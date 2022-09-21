package tools

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
)

func SubchainChannelRegister(targetChainID *big.Int, IP string, sourceChainEthRpcClientURL string) {
	subchainClient, err := ethclient.Dial(sourceChainEthRpcClientURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for RegisterChannelOnSubchain for chainID %v and IP is %v...\n", targetChainID.String(), IP)
	subchainRegisterAddr := common.HexToAddress("0xBd770416a3345F91E4B34576cb804a576fa48EB1")
	subchainRegisterInstance, _ := ct.NewChainRegistrarOnSubchain(subchainRegisterAddr, subchainClient)

	authUser := subchainSelectAccount(subchainClient, 1)
	regitserTx, err := subchainRegisterInstance.RegisterSubchainChannel(authUser, targetChainID, IP)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("registerTX", regitserTx.Hash().Hex())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), regitserTx.Hash())
	time.Sleep(6 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("register error")
	}

}

func GetMaxProcessedNonceFromRegistrar() {
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for QueryMaxNonceForSubchainChannel...\n")
	subchainRegisterAddr := common.HexToAddress("0xBd770416a3345F91E4B34576cb804a576fa48EB1")
	subchainRegisterInstance, _ := ct.NewChainRegistrarOnSubchain(subchainRegisterAddr, subchainClient)

	maxProcessedSubchainRegisteredNonce, err := subchainRegisterInstance.GetMaxProcessedNonce(nil)
	if err != nil {
		log.Fatal(err)
		return // ignore
	}
	log.Println("Max nonce : ", maxProcessedSubchainRegisteredNonce)

}

func GetCrossChainFeeFromRegistrar() {
	subchainClient, err := ethclient.Dial("http://localhost:19888/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for Chain Fee...\n")
	subchainRegisterAddr := common.HexToAddress("0xBd770416a3345F91E4B34576cb804a576fa48EB1")
	subchainRegisterInstance, _ := ct.NewChainRegistrarOnSubchain(subchainRegisterAddr, subchainClient)

	maxProcessedSubchainRegisteredNonce, err := subchainRegisterInstance.GetCrossChainFee(nil)
	if err != nil {
		log.Fatal(err)
		return // ignore
	}
	log.Println("Chain Fee : ", maxProcessedSubchainRegisteredNonce)

}

func GetChannelStatusFromRegistrar(targetChainID *big.Int, targetChainEthRpcClientURL string) {
	subchainClient, err := ethclient.Dial(targetChainEthRpcClientURL)
	if err != nil {
		log.Fatal(err)
	}
	localChainID, err := subchainClient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for Query Channel Status...\n")
	subchainRegisterAddr := common.HexToAddress("0xBd770416a3345F91E4B34576cb804a576fa48EB1")
	subchainRegisterInstance, _ := ct.NewChainRegistrarOnSubchain(subchainRegisterAddr, subchainClient)

	channelStatus, err := subchainRegisterInstance.IsAnActiveChannel(nil, targetChainID)
	if err != nil {
		log.Fatal(err)
		return // ignore
	}
	log.Printf("The Channel from %v to %v is active? : %v", localChainID.String(), targetChainID.String(), channelStatus)
	dynasty, _, err := subchainRegisterInstance.GetDynasty(nil)
	var a struct {
		Validators   []common.Address
		ShareAmounts []*big.Int
	}
	a, _ = subchainRegisterInstance.GetValidatorSet(nil, targetChainID, dynasty)
	fmt.Println(a.Validators)
}

func VerifyChannel(targetChainID *big.Int, targetChainEthRpcClientURL string) {
	subchainClient, err := ethclient.Dial(targetChainEthRpcClientURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for Verify Query Channel for chainID %v, IP is %v...\n", targetChainID.String(), targetChainEthRpcClientURL)
	subchainRegisterAddr := common.HexToAddress("0xBd770416a3345F91E4B34576cb804a576fa48EB1")
	subchainRegisterInstance, _ := ct.NewChainRegistrarOnSubchain(subchainRegisterAddr, subchainClient)
	authUser := subchainSelectAccount(subchainClient, 1)
	fmt.Println(authUser.GasPrice)
	tx, err := subchainRegisterInstance.UpdateSubchainChannelStatus(authUser, targetChainID, true, big.NewInt(2))
	if err != nil {
		log.Fatal(err)
		return // ignore
	}
	fmt.Println(tx.Hash().Hex())
}
func Testlock() {
	lockAmount := big.NewInt(10)
	subchainClient, err := ethclient.Dial("http://localhost:19988/rpc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Preparing for TNT20 cross-chain transfer...\n")

	subchainTNT20Address := common.HexToAddress("0x5C3159dDD2fe0F9862bC7b7D60C1875fa8F81337")

	sender := accountList[1].fromAddress
	//receiver := accountList[6].fromAddress
	subchainTNT20TokenBankInstance, _ := ct.NewTNT20TokenBank(subchainTNT20TokenBankAddress, subchainClient)
	subchainTNT20Instance, _ := ct.NewMockTNT20(subchainTNT20Address, subchainClient)

	mintAmount := big.NewInt(1).Mul(big.NewInt(5), lockAmount)
	fmt.Printf("Minting %v TNT20 tokens\n", mintAmount)

	authUser := subchainSelectAccount(subchainClient, 1)
	subchainTNT20Instance.Mint(authUser, sender, mintAmount)
	time.Sleep(6 * time.Second)

	senderTNT20Balance, _ := subchainTNT20Instance.BalanceOf(nil, sender)
	subchainTNT20Name, _ := subchainTNT20Instance.Name(nil)
	subchainTNT20Symbol, _ := subchainTNT20Instance.Symbol(nil)
	subchainTNT20Decimals, _ := subchainTNT20Instance.Decimals(nil)

	fmt.Printf("Subchain TNT20 contract address: %v, Name: %v, Symbol: %v, Decimals: %v\n", subchainTNT20Address, subchainTNT20Name, subchainTNT20Symbol, subchainTNT20Decimals)
	fmt.Printf("Subchain sender   : %v, TNT20 balance on Subchain         : %v\n", sender, senderTNT20Balance)

	authUser = subchainSelectAccount(subchainClient, 1)
	subchainTNT20Instance.Approve(authUser, subchainTNT20TokenBankAddress, lockAmount)

	authUser = subchainSelectAccount(subchainClient, 1)
	authUser.Value.Set(crossChainFee)
	lockTx, err := subchainTNT20TokenBankInstance.LockTokens(authUser, big.NewInt(360777), subchainTNT20Address, accountList[6].fromAddress, lockAmount)
	authUser.Value.Set(common.Big0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TNT20 Token Lock tx hash (Subchain): %v\n", lockTx.Hash().Hex())
	fmt.Printf("Transfering %v TNT20 tokens (Wei) from to Subchain %v to the Mainchain...\n\n", lockAmount, subchainID)

	fmt.Printf("Start transfer, timestamp      : %v\n", time.Now())
	receipt, err := subchainClient.TransactionReceipt(context.Background(), lockTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != 1 {
		log.Fatal("lock error")
	}
	fmt.Printf("Token lock confirmed, timestamp: %v\n", time.Now())

}
