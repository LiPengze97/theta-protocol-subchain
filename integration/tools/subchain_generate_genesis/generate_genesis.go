package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/rlp"
	"github.com/thetatoken/theta/store/database/backend"
	"github.com/thetatoken/theta/store/trie"

	scom "github.com/thetatoken/thetasubchain/common"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/interchain/contracts/predeployed"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	svm "github.com/thetatoken/thetasubchain/ledger/vm"
)

var logger *log.Entry = log.WithFields(log.Fields{"prefix": "genesis"})

const (
	GenBlockHashMode int = iota
	GenGenesisFileMode
)

type Validator struct {
	Address string `json:"address"`
	Stake   string `json:"stake"`
}

//
// Example:
// cd $SUBCHAIN_HOME/integration/privatenet/node
// subchain_generate_genesis -mainchainID=privatenet -subchainID=tsub_360777 -initValidatorSet=./data/init_validator_set.json -genesis=./genesis
//
func main() {
	mainchainID, subchainID, initValidatorSetPath, genesisSnapshotFilePath := parseArguments()

	sv, metadata, err := generateGenesisSnapshot(mainchainID, subchainID, initValidatorSetPath, genesisSnapshotFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate genesis snapshot: %v", err))
	}

	err = sanityChecks(sv)
	if err != nil {
		panic(fmt.Sprintf("Sanity checks failed: %v", err))
	} else {
		logger.Infof("Sanity checks all passed.")
	}

	err = writeGenesisSnapshot(sv, metadata, genesisSnapshotFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to write genesis snapshot: %v", err))
	}

	genesisBlockHeader := metadata.TailTrio.Second.Header
	genesisBlockHash := genesisBlockHeader.Hash()

	fmt.Println("")
	fmt.Printf("-----------------------------------------------------------------------------------------\n")
	fmt.Printf("Genesis block hash: %v\n", genesisBlockHash.Hex())
	fmt.Printf("-----------------------------------------------------------------------------------------\n")
	fmt.Println("")
}

func parseArguments() (mainchainID, subchainID, initValidatorSetPath, genesisSnapshotFilePath string) {
	mainchainIDPtr := flag.String("mainchainID", "privatenet", "the ID of the mainchain")
	subchainIDPtr := flag.String("subchainID", "tsub_360777", "the ID of the subchain")
	initValidatorSetPathPtr := flag.String("initValidatorSet", "./init_validator_set.json", "the initial validator set")
	genesisSnapshotFilePathPtr := flag.String("genesis", "./genesis", "the genesis snapshot")
	flag.Parse()

	mainchainID = *mainchainIDPtr
	subchainID = *subchainIDPtr
	initValidatorSetPath = *initValidatorSetPathPtr
	genesisSnapshotFilePath = *genesisSnapshotFilePathPtr

	return
}

// generateGenesisSnapshot generates the genesis snapshot.
func generateGenesisSnapshot(mainchainID, subchainID, initValidatorSetFilePath, genesisSnapshotFilePath string) (*slst.StoreView, *score.SnapshotMetadata, error) {
	metadata := &score.SnapshotMetadata{}
	genesisHeight := score.GenesisBlockHeight

	sv := slst.NewStoreView(0, common.Hash{}, backend.NewMemDatabase())

	setInitialValidatorSet(subchainID, initValidatorSetFilePath, genesisHeight, sv)
	deployInitialSmartContracts(mainchainID, subchainID, sv)

	stateHash := sv.Hash()

	genesisBlock := score.NewBlock()
	genesisBlock.ChainID = subchainID
	genesisBlock.Height = genesisHeight
	genesisBlock.Epoch = genesisBlock.Height
	genesisBlock.Parent = common.Hash{}
	genesisBlock.StateHash = stateHash
	genesisBlock.Timestamp = big.NewInt(time.Now().Unix())

	metadata.TailTrio = score.SnapshotBlockTrio{
		First:  score.SnapshotFirstBlock{},
		Second: score.SnapshotSecondBlock{Header: genesisBlock.BlockHeader},
		Third:  score.SnapshotThirdBlock{},
	}

	return sv, metadata, nil
}

func setInitialValidatorSet(subchainID string, initValidatorSetFilePath string, genesisHeight uint64, sv *slst.StoreView) *score.ValidatorSet {
	var validators []Validator
	initValidatorSetFile, err := os.Open(initValidatorSetFilePath)
	if err != nil {
		panic(fmt.Sprintf("failed to open initial stake deposit file: %v", err))
	}
	initValidatorSetByteValue, err := ioutil.ReadAll(initValidatorSetFile)
	if err != nil {
		panic(fmt.Sprintf("failed to read initial stake deposit file: %v", err))
	}

	json.Unmarshal(initValidatorSetByteValue, &validators)
	initialDynasty := big.NewInt(0)
	validatorSet := score.NewValidatorSet(initialDynasty)
	for _, v := range validators {
		stake, ok := big.NewInt(0).SetString(v.Stake, 10)
		if !ok {
			panic(fmt.Sprintf("failed to parse stake: %v", v.Stake))
		}
		validator := score.NewValidator(v.Address, stake)
		validatorSet.AddValidator(validator)

		setInitialBalance(sv, common.HexToAddress(v.Address), big.NewInt(0)) // need to create an account with zero balance for the initial validators
	}

	subchainIDInt := scom.MapChainID(subchainID)
	sv.UpdateValidatorSet(subchainIDInt, validatorSet)

	hl := &types.HeightList{}
	hl.Append(genesisHeight)
	sv.UpdateValidatorSetUpdateTxHeightList(hl)
	sv.Save()

	return validatorSet
}

func setInitialBalance(sv *slst.StoreView, address common.Address, tfuelBalance *big.Int) {
	acc := &types.Account{
		Address:  address,
		Root:     common.Hash{},
		CodeHash: types.EmptyCodeHash,
		Balance: types.Coins{
			ThetaWei: big.NewInt(0),
			TFuelWei: tfuelBalance,
		},
	}
	sv.SetAccount(acc.Address, acc)
}

func proveValidatorSet(sv *slst.StoreView) (*score.ValidatorSetProof, error) {
	vp := &score.ValidatorSetProof{}
	vsKey := slst.CurrentValidatorSetKey()
	err := sv.ProveValidatorSet(vsKey, vp)
	return vp, err
}

func deployInitialSmartContracts(mainchainID, subchainID string, sv *slst.StoreView) {
	mainchainIDInt := scom.MapChainID(mainchainID)
	deployer := common.Address{}

	//
	// Deploy the ChainRegistrar contract
	//

	sequence := 0
	chainRegistrarContractAddr, err := deploySmartContract(subchainID, sv, predeployed.ChainRegistrarContractBytecode, deployer, sequence, slst.ChainRegistrarContractAddressKey())
	if err != nil {
		logger.Panicf("Failed to deploy the chain registrar smart contract (sequence = %v): %v", sequence, err)
	}

	tokenBankContractBytecodes := []string{
		predeployed.TFuelTokenBankContractBytecode,
		predeployed.TNT20TokenBankContractBytecode,
		//predeployed.TNT721TokenBankContractBytecode,
	}
	contractAddressKeys := []common.Bytes{
		slst.TFuelTokenBankContractAddressKey(),
		slst.TNT20TokenBankContractAddressKey(),
		//slst.TNT721TokenBankContractAddressKey(),
	}

	//
	// Deploy the TokenBank contracts
	//

	for idx, contractBytecode := range tokenBankContractBytecodes {
		sequence += 1
		tokenBankDeploymentBytecode := addConstructorArgumentForTokenBankBytecode(contractBytecode, mainchainIDInt, chainRegistrarContractAddr)
		contractAddressKey := contractAddressKeys[idx]
		_, err := deploySmartContract(subchainID, sv, tokenBankDeploymentBytecode, deployer, sequence, contractAddressKey)
		if err != nil {
			logger.Panicf("Failed to deploy TokenBank smart contract (sequence = %v): %v", sequence, err)
		}
	}
}

// Reference: https://docs.blockscout.com/for-users/abi-encoded-constructor-arguments
func addConstructorArgumentForTokenBankBytecode(contractBytecode string, mainchainIDInt *big.Int, chainRegistrarContractAddr common.Address) string {
	// TODO: implementation
	return ""
}

func deploySmartContract(subchainID string, sv *slst.StoreView, contractBytecodeStr string, deployer common.Address, sequence int, contractAddressKey common.Bytes) (common.Address, error) {
	dummyGasLimit := uint64(10000000)
	dummyGasPrice := big.NewInt(1)

	// Token Bank contract
	contractBytecode, err := hex.DecodeString(contractBytecodeStr)
	if err != nil {
		return common.Address{}, err
	}
	deploySCTx := types.SmartContractTx{
		From:     types.NewTxInput(common.Address{}, types.NewCoins(0, 0), sequence),
		To:       types.TxOutput{Address: common.Address{}}, // deploy contract
		GasLimit: dummyGasLimit,
		GasPrice: dummyGasPrice,
		Data:     contractBytecode,
	}
	parentBlockInfo := svm.NewBlockInfo(0, big.NewInt(0), subchainID)
	_, contractAddr, _, evmErr := svm.Execute(parentBlockInfo, &deploySCTx, sv)
	if evmErr != nil {
		return common.Address{}, evmErr
	}

	tbcaBytes, err := types.ToBytes(contractAddr)
	if err != nil {
		return common.Address{}, err
	}
	sv.Set(contractAddressKey, tbcaBytes)

	return contractAddr, nil
}

// writeGenesisSnapshot writes genesis snapshot to file system.
func writeGenesisSnapshot(sv *slst.StoreView, metadata *score.SnapshotMetadata, genesisSnapshotFilePath string) error {
	file, err := os.Create(genesisSnapshotFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	err = score.WriteMetadata(writer, metadata)
	if err != nil {
		return err
	}
	writeStoreView(sv, true, writer)
	return err
}

func writeStoreView(sv *slst.StoreView, needAccountStorage bool, writer *bufio.Writer) {
	height := score.Itobytes(sv.Height())
	err := score.WriteRecord(writer, []byte{score.SVStart}, height)
	if err != nil {
		panic(err)
	}
	sv.GetStore().Traverse(nil, func(k, v common.Bytes) bool {
		err = score.WriteRecord(writer, k, v)
		if err != nil {
			panic(err)
		}
		return true
	})
	err = score.WriteRecord(writer, []byte{score.SVEnd}, height)
	if err != nil {
		panic(err)
	}
	writer.Flush()
}

func sanityChecks(sv *slst.StoreView) error {
	logger.Infof("------------------------------------------------------------------------------")

	vsAnalyzed := false
	sv.GetStore().Traverse(nil, func(key, val common.Bytes) bool {
		if bytes.Equal(key, slst.CurrentValidatorSetKey()) {
			var vs score.ValidatorSet
			err := rlp.DecodeBytes(val, &vs)
			if err != nil {
				panic(fmt.Sprintf("Failed to decode validator set: %v", err))
			}
			for _, validator := range vs.Validators() {
				logger.Infof("Initial validator: %v, stake: %v", validator.Address.Hex(), validator.Stake)
			}
			vsAnalyzed = true
		} else if bytes.Equal(key, slst.ValidatorSetUpdateTxHeightListKey()) {
			var hl types.HeightList
			err := rlp.DecodeBytes(val, &hl)
			if err != nil {
				panic(fmt.Sprintf("Failed to decode Height List: %v", err))
			}
			if len(hl.Heights) != 1 {
				panic(fmt.Sprintf("The genesis height list should contain only one height: %v", hl.Heights))
			}
			if hl.Heights[0] != uint64(0) {
				panic("Only height 0 should be in the genesis height list")
			}
		}
		return true
	})

	tfuelTokenBankContractAddr := sv.GetTFuelTokenBankContractAddress()
	if tfuelTokenBankContractAddr == nil {
		panic("TFuel token bank contract is not set")
	}
	logger.Infof("TFuel Token Bank Contract Address: %v", tfuelTokenBankContractAddr.Hex())

	tnt20TokenBankContractAddr := sv.GetTNT20TokenBankContractAddress()
	if tnt20TokenBankContractAddr == nil {
		panic("TNT20 token bank contract is not set")
	}
	logger.Infof("TNT20 Token Bank Contract Address: %v", tnt20TokenBankContractAddr.Hex())

	// tnt721TokenBankContractAddr := sv.GetTNT721TokenBankContractAddress()
	// if tnt721TokenBankContractAddr == nil {
	// 	panic("TNT721 token bank contract is not set")
	// }
	// logger.Infof("TNT721 Token Bank Contract Address: %v", tnt721TokenBankContractAddr.Hex())

	logger.Infof("------------------------------------------------------------------------------")

	// Sanity checks for the initial validator set
	vsProof, err := proveValidatorSet(sv)
	if err != nil {
		panic("Failed to get VS proof from storeview")
	}
	_, _, err = trie.VerifyProof(sv.Hash(), slst.CurrentValidatorSetKey(), vsProof)
	if err != nil {
		panic("Failed to verify VS proof in storeview")
	}
	if !vsAnalyzed {
		return fmt.Errorf("initial validator set not detected in the genesis file")
	}

	return nil
}
