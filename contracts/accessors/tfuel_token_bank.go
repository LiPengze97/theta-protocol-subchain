// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accessors

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/thetatoken/thetasubchain/eth"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	"github.com/thetatoken/thetasubchain/eth/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TFuelTokenBankMetaData contains all meta data concerning the TFuelTokenBank contract.
var TFuelTokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"mainchainid_\",\"type\":\"uint256\"},{\"internaltype\":\"contractsubchainregistrar\",\"name\":\"subchainregistrar_\",\"type\":\"address\"}],\"statemutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"sourcechaintokensender\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"targetchainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"targetchainvoucherreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"lockedamount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"tokenlocknonce\",\"type\":\"uint256\"}],\"name\":\"tfueltokenlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"targetchaintokenreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"unlockedamount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"sourcechainvoucherburnnonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"tokenunlocknonce\",\"type\":\"uint256\"}],\"name\":\"tfueltokenunlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"sourcechainvoucherowner\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"targetchaintokenreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"burnedamount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"voucherburnnonce\",\"type\":\"uint256\"}],\"name\":\"tfuelvoucherburned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"targetchainvoucherreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"mintedamount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"sourcechaintokenlocknonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"vouchermintnonce\",\"type\":\"uint256\"}],\"name\":\"tfuelvoucherminted\",\"type\":\"event\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alldenoms\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allvouchers\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomtovoucherlookup\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"contractaddress\",\"type\":\"address\"},{\"internaltype\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"voucheraddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"vouchercontractaddr\",\"type\":\"address\"}],\"name\":\"getdenom\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"getmaxprocessedtokenlocknonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"getmaxprocessedvoucherburnnonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getvoucher\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isonmainchain\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainchainid\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenlocknoncemap\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internaltype\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenlockvotingrecords\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"accumlatedshares\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenunlocknoncemap\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totallockedamounts\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucheraddresstodenomlookup\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internaltype\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherburnnoncemap\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internaltype\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherburnvotingrecords\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"accumlatedshares\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vouchermintnoncemap\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"targetchainid\",\"type\":\"uint256\"},{\"internaltype\":\"address\",\"name\":\"targetchainvoucherreceiver\",\"type\":\"address\"}],\"name\":\"locktokens\",\"outputs\":[],\"statemutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internaltype\":\"address\",\"name\":\"targetchainvoucherreceiver\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"mintedamount\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"sourcechaintokenlocknonce\",\"type\":\"uint256\"}],\"name\":\"mintvouchers\",\"outputs\":[],\"statemutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"targetchaintokenreceiver\",\"type\":\"address\"}],\"name\":\"burnvouchers\",\"outputs\":[],\"statemutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"sourcechainid\",\"type\":\"uint256\"},{\"internaltype\":\"addresspayable\",\"name\":\"targetchaintokenreceiver\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"unlockamount\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"sourcechainvoucherburnnonce\",\"type\":\"uint256\"}],\"name\":\"unlocktokens\",\"outputs\":[],\"statemutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200262f3803806200262f833981016040819052620000349162000064565b6001600081905591909155600280546001600160a01b0319166001600160a01b03909216919091179055620000a3565b600080604083850312156200007857600080fd5b825160208401519092506001600160a01b03811681146200009857600080fd5b809150509250929050565b61257c80620000b36000396000f3fe6080604052600436106101405760003560e01c8063766f8fb0116100b6578063ccf187c71161006f578063ccf187c714610454578063ebda996214610481578063f6a3d24e146104a1578063f95627ac146104dd578063feaff0521461050a578063ff248a441461054957600080fd5b8063766f8fb0146103875780637d0fb00d146103b45780638883931e146103c7578063a2cc6981146103f4578063aa68acde14610414578063ca2075691461042757600080fd5b8063261a323e11610108578063261a323e1461027d57806327ca4df1146102ad5780634250863b146102e5578063588b1408146102ff57806360569b5e1461032c578063740cb7f81461035a57600080fd5b8063073b9502146101455780631527b14d1461016e57806319fd1a11146101da5780631a0483d3146102075780631eb7873714610229575b600080fd5b34801561015157600080fd5b5061015b60015481565b6040519081526020015b60405180910390f35b34801561017a57600080fd5b506101bb610189366004611fba565b8051602081830181018051600b825292820191909301209152546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b039093168352901515602083015201610165565b3480156101e657600080fd5b5061015b6101f5366004612059565b60106020526000908152604090205481565b34801561021357600080fd5b50610227610222366004611fef565b61055c565b005b34801561023557600080fd5b506102686102443660046120cc565b60096020908152600092835260408084209091529082529020805460029091015482565b60408051928352602083019190915201610165565b34801561028957600080fd5b5061029d610298366004611fba565b61073e565b6040519015158152602001610165565b3480156102b957600080fd5b506102cd6102c8366004612059565b61077e565b6040516001600160a01b039091168152602001610165565b3480156102f157600080fd5b50600f5461029d9060ff1681565b34801561030b57600080fd5b5061031f61031a366004612059565b6107a8565b60405161016591906121e9565b34801561033857600080fd5b5061034c610347366004611eaf565b610854565b6040516101659291906122c2565b34801561036657600080fd5b5061015b610375366004612059565b60066020526000908152604090205481565b34801561039357600080fd5b5061015b6103a2366004612059565b60009081526008602052604090205490565b6102276103c2366004611eaf565b6108fb565b3480156103d357600080fd5b5061015b6103e2366004612059565b60036020526000908152604090205481565b34801561040057600080fd5b506102cd61040f366004611fba565b610a50565b610227610422366004612072565b610ac1565b34801561043357600080fd5b5061015b610442366004612059565b60056020526000908152604090205481565b34801561046057600080fd5b5061015b61046f366004612059565b60046020526000908152604090205481565b34801561048d57600080fd5b5061031f61049c366004611eaf565b610d02565b3480156104ad57600080fd5b5061029d6104bc366004611eaf565b6001600160a01b03166000908152600c602052604090206001015460ff1690565b3480156104e957600080fd5b5061015b6104f8366004612059565b60009081526007602052604090205490565b34801561051657600080fd5b506102686105253660046120cc565b600a6020908152600092835260408084209091529082529020805460029091015482565b6102276105573660046120a2565b610df9565b600260005414156105885760405162461bcd60e51b815260040161057f906122e6565b60405180910390fd5b60026000556001544614156106055760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6d696e74566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e0000606482015260840161057f565b600061061086611092565b905060008061061e836110a3565b915091508061067b5760405162461bcd60e51b8152602060048201526024808201527f4661696c656420746f206578747261637420636861696e49442066726f6d2064604482015263656e6f6d60e01b606482015260840161057f565b600082848989898960405160200161069896959493929190612190565b60405160208183030381529060405280519060200120905060006106bf8488848933611134565b9050801561072d576106d18989611151565b6106da46611330565b46600090815260066020526040908190205490517f80742bd15a2c8c4ad5d395bcf577073110e52f0c73bf980dfa9453c1d8c354e590610723908d908d908d908c9087906121fc565b60405180910390a1505b505060016000555050505050505050565b60008061074a83611092565b9050600b8160405161075c919061211a565b9081526040519081900360200190205460ff600160a01b909104169392505050565b600d818154811061078e57600080fd5b6000918252602090912001546001600160a01b0316905081565b600e81815481106107b857600080fd5b9060005260206000200160009150905080546107d390612448565b80601f01602080910402602001604051908101604052809291908181526020018280546107ff90612448565b801561084c5780601f106108215761010080835404028352916020019161084c565b820191906000526020600020905b81548152906001019060200180831161082f57829003601f168201915b505050505081565b600c6020526000908152604090208054819061086f90612448565b80601f016020809104026020016040519081016040528092919081815260200182805461089b90612448565b80156108e85780601f106108bd576101008083540402835291602001916108e8565b820191906000526020600020905b8154815290600101906020018083116108cb57829003601f168201915b5050506001909301549192505060ff1682565b6002600054141561091e5760405162461bcd60e51b815260040161057f906122e6565b600260005560015446141561099b5760405162461bcd60e51b815260206004820152603e60248201527f544675656c546f6b656e42616e6b2e6275726e566f756368657273282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e206120737562636861696e0000606482015260840161057f565b33346109a681611357565b6109af466114bb565b60006109ef600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a815260200161251d602a91396114da565b4660009081526005602052604090819020549051919250907f40f1d475c2aa44f5c23193fab26a64d6aa4e09ab51898b10a3036baf82398ea190610a3c908490879089908890879061223c565b60405180910390a150506001600055505050565b600080610a5c83611092565b90506000600b82604051610a70919061211a565b908152604080516020928190038301812081830190925290546001600160a01b0381168252600160a01b900460ff16158015928201929092529150610ab757519392505050565b5060009392505050565b60026000541415610ae45760405162461bcd60e51b815260040161057f906122e6565b60026000556001544614610b62576040805162461bcd60e51b81526020600482015260248101919091527f544675656c546f6b656e42616e6b2e6c6f636b546f6b656e7328292063616e2060448201527f6f6e6c792062652063616c6c6564206f6e20746865206d61696e20636861696e606482015260840161057f565b6002546040516343b71f0560e01b8152600481018490526001600160a01b03909116906343b71f059060240160206040518083038186803b158015610ba657600080fd5b505afa158015610bba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bde9190611f98565b610c2a5760405162461bcd60e51b815260206004820181905260248201527f746172676574436861696e4944206e6f74207965742072656769737465726564604482015260640161057f565b60008281526010602052604090205434903390610c479083611518565b600085815260106020526040902055610c5f8461152b565b6000610c9f600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a815260200161251d602a91396114da565b60008681526003602052604090819020549051919250907fee1ecc2b21aa613cc77cd44823a68ef1168ce1f40c2eac1d68690baf955fdbd190610ced90849086908a908a908a90889061227c565b60405180910390a15050600160005550505050565b6001600160a01b0381166000908152600c60205260408082208151808301909252805460609392919082908290610d3890612448565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6490612448565b8015610db15780601f10610d8657610100808354040283529160200191610db1565b820191906000526020600020905b815481529060010190602001808311610d9457829003601f168201915b50505091835250506001919091015460ff16151560209182015281015190915015610ddd575192915050565b5050604080516020810190915260008152919050565b50919050565b60026000541415610e1c5760405162461bcd60e51b815260040161057f906122e6565b60026000556001544614610ea35760405162461bcd60e51b815260206004820152604260248201527f544675656c546f6b656e42616e6b2e756e6c6f636b546f6b656e73282920636160448201527f6e206f6e6c792062652063616c6c6564206f6e20746865206d61696e2063686160648201526134b760f11b608482015260a40161057f565b600085815260106020526040902054831115610f155760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f7420756e6c6f636b207468652072657175657374656420616d6f7560448201526a1b9d081bd98815119d595b60aa1b606482015260840161057f565b6040805160208101879052908101849052606085811b6bffffffffffffffffffffffff191690820152607481018390526094810182905260009060b40160408051601f19818403018152919052805160208201209091506000610f7b888684873361154a565b9050801561108357600088815260106020526040902054610f9c908761155d565b6000898152601060205260408082209290925590516001600160a01b0389169188156108fc02918991818181858888f19350505050158015610fe2573d6000803e3d6000fd5b50610fec88611569565b600061102c600154604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a815260200161251d602a91396114da565b60008a81526004602052604090819020549051919250907f5ea3a5ca7f54881fdd7781894d69709e11027910f35647f9d4cc14e6872b6f72906110789084908c908c908b9087906121fc565b60405180910390a150505b50506001600055505050505050565b606061109d82611588565b92915050565b6000806000806110b985602f60f81b6001611602565b91509150806110cf575060009485945092505050565b6000806110e787826110e2876001612372565b611692565b91509150806110ff5750600096879650945050505050565b60008061110b846117a9565b915091508061112557506000988998509650505050505050565b50976001975095505050505050565b600061114786868686866007600961187b565b9695505050505050565b60408051603480825260608201909252600091602082018180368337019050509050816000805b60148110156111e0578560601b8160148110611196576111966124d8565b1a60f81b84836111a58161247d565b9450815181106111b7576111b76124d8565b60200101906001600160f81b031916908160001a905350806111d88161247d565b915050611178565b5060005b6020811015611249578281602081106111ff576111ff6124d8565b1a60f81b848361120e8161247d565b945081518110611220576112206124d8565b60200101906001600160f81b031916908160001a905350806112418161247d565b9150506111e4565b600060b66001600160a01b031685604051611264919061211a565b6000604051808303816000865af19150503d80600081146112a1576040519150601f19603f3d011682016040523d82523d6000602084013e6112a6565b606091505b50509050806113275760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6d696e74544675656c566f756368657260448201527f28293a206661696c656420746f206d696e7420544675656c20766f75636865726064820152607360f81b608482015260a40161057f565b50505050505050565b600081815260066020526040812080546001929061134f908490612372565b909155505050565b6040805160208082528183019092526000916020820181803683370190505090508160005b60208110156113d657818160208110611397576113976124d8565b1a60f81b8382815181106113ad576113ad6124d8565b60200101906001600160f81b031916908160001a905350806113ce8161247d565b91505061137c565b50600060b76001600160a01b0316836040516113f2919061211a565b6000604051808303816000865af19150503d806000811461142f576040519150601f19603f3d011682016040523d82523d6000602084013e611434565b606091505b50509050806114b55760405162461bcd60e51b815260206004820152604160248201527f544675656c546f6b656e42616e6b2e5f6275726e544675656c566f756368657260448201527f28293a206661696c656420746f206275726e20544675656c20766f75636865726064820152607360f81b608482015260a40161057f565b50505050565b600081815260056020526040812080546001929061134f908490612372565b60606115106114e885611c74565b84846040516020016114fc93929190612136565b604051602081830303815290604052611092565b949350505050565b60006115248284612372565b9392505050565b600081815260036020526040812080546001929061134f908490612372565b600061114786868686866008600a61187b565b600061152482846123e2565b600081815260046020526040812080546001929061134f908490612372565b60608160005b81518110156115fb576115c08282815181106115ac576115ac6124d8565b01602001516001600160f81b031916611d72565b8282815181106115d2576115d26124d8565b60200101906001600160f81b031916908160001a905350806115f38161247d565b91505061158e565b5092915050565b82516000908190859082805b8281101561167e57876001600160f81b031916848281518110611633576116336124d8565b01602001516001600160f81b031916141561166c57611653600183612372565b91508682141561166c5794506001935061168a92505050565b806116768161247d565b91505061160e565b50600080945094505050505b935093915050565b8251606090600090848410156116bc5750506040805160208101909152600080825290915061168a565b6116c781600161155d565b8411156116e85750506040805160208101909152600080825290915061168a565b8560006116f5868861155d565b67ffffffffffffffff81111561170d5761170d6124ee565b6040519080825280601f01601f191660200182016040528015611737576020820181803683370190505b509050865b8681101561112557828181518110611756576117566124d8565b01602001516001600160f81b03191682611770838b61155d565b81518110611780576117806124d8565b60200101906001600160f81b031916908160001a905350806117a18161247d565b91505061173c565b80516000908190839082805b8281101561186e5760308482815181106117d1576117d16124d8565b016020015160f81c10801590611801575060398482815181106117f6576117f66124d8565b016020015160f81c11155b1561184d5761181182600a611dc1565b9150611846603085838151811061182a5761182a6124d8565b016020015161183c919060f81c6123f9565b839060ff16611518565b915061185c565b50600096879650945050505050565b806118668161247d565b9150506117b5565b5095600195509350505050565b600060015488148061190457506002546040516343b71f0560e01b8152600481018a90526001600160a01b03909116906343b71f059060240160206040518083038186803b1580156118cc57600080fd5b505afa1580156118e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119049190611f98565b6119425760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9251608a1b604482015260640161057f565b60008881526020849052604090205461195c906001612372565b851461196a57506000611c69565b600088815260208381526040808320898452909152812060015460609081908c908114156119955750465b6002546040516343f27e4560e01b815260048101839052602481018e90526001600160a01b03909116906343f27e459060440160006040518083038186803b1580156119e057600080fd5b505afa1580156119f4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a1c9190810190611ecc565b9350915060005b8251811015611b7957896001600160a01b0316838281518110611a4857611a486124d8565b60200260200101516001600160a01b031614611a6357611b67565b6001955060005b6001860154811015611b0457856001018181548110611a8b57611a8b6124d8565b6000918252602090912001546001600160a01b038c811691161415611af25760405162461bcd60e51b815260206004820152601c60248201527f546869732076616c696461746f7220616c726561647920766f74656400000000604482015260640161057f565b80611afc8161247d565b915050611a6a565b508c85556001808601805491820181556000908152602090200180546001600160a01b031916331790558351611b6190859083908110611b4657611b466124d8565b6020026020010151866002015461151890919063ffffffff16565b60028601555b80611b718161247d565b915050611a23565b50505082611bbb5760405162461bcd60e51b815260206004820152600f60248201526e2737ba1030903b30b634b230ba37b960891b604482015260640161057f565b6000805b8251811015611c0857611bf4838281518110611bdd57611bdd6124d8565b60200260200101518361151890919063ffffffff16565b915080611c008161247d565b915050611bbf565b50611c14816002611dc1565b6002840154611c24906003611dc1565b10611c605760008c815260208890526040902054611c43906001612372565b60008d8152602089905260409020555060019350611c6992505050565b60009450505050505b979650505050505050565b606081611c985750506040805180820190915260018152600360fc1b602082015290565b8160005b8115611cc25780611cac8161247d565b9150611cbb9050600a836123af565b9150611c9c565b60008167ffffffffffffffff811115611cdd57611cdd6124ee565b6040519080825280601f01601f191660200182016040528015611d07576020820181803683370190505b5090505b841561151057611d1c6001836123e2565b9150611d29600a86612498565b611d34906030612372565b60f81b818381518110611d4957611d496124d8565b60200101906001600160f81b031916908160001a905350611d6b600a866123af565b9450611d0b565b6000604160f81b6001600160f81b0319831610801590611da05750602d60f91b6001600160f81b0319831611155b15611dbd57611db460f883901c602061238a565b60f81b92915050565b5090565b600061152482846123c3565b600082601f830112611dde57600080fd5b81516020611df3611dee8361234e565b61231d565b80838252828201915082860187848660051b8901011115611e1357600080fd5b60005b85811015611e3257815184529284019290840190600101611e16565b5090979650505050505050565b600082601f830112611e5057600080fd5b813567ffffffffffffffff811115611e6a57611e6a6124ee565b611e7d601f8201601f191660200161231d565b818152846020838601011115611e9257600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611ec157600080fd5b813561152481612504565b60008060408385031215611edf57600080fd5b825167ffffffffffffffff80821115611ef757600080fd5b818501915085601f830112611f0b57600080fd5b81516020611f1b611dee8361234e565b8083825282820191508286018a848660051b8901011115611f3b57600080fd5b600096505b84871015611f67578051611f5381612504565b835260019690960195918301918301611f40565b5091880151919650909350505080821115611f8157600080fd5b50611f8e85828601611dcd565b9150509250929050565b600060208284031215611faa57600080fd5b8151801515811461152457600080fd5b600060208284031215611fcc57600080fd5b813567ffffffffffffffff811115611fe357600080fd5b61151084828501611e3f565b600080600080600060a0868803121561200757600080fd5b853567ffffffffffffffff81111561201e57600080fd5b61202a88828901611e3f565b955050602086013561203b81612504565b94979496505050506040830135926060810135926080909101359150565b60006020828403121561206b57600080fd5b5035919050565b6000806040838503121561208557600080fd5b82359150602083013561209781612504565b809150509250929050565b600080600080600060a086880312156120ba57600080fd5b85359450602086013561203b81612504565b600080604083850312156120df57600080fd5b50508035926020909101359150565b6000815180845261210681602086016020860161241c565b601f01601f19169290920160200192915050565b6000825161212c81846020870161241c565b9190910192915050565b6000845161214881846020890161241c565b8083019050602f60f81b8082528551612168816001850160208a0161241c565b6001920191820152835161218381600284016020880161241c565b0160020195945050505050565b868152600086516121a8816020850160208b0161241c565b60609690961b6bffffffffffffffffffffffff1916602092909601918201959095526034810193909352605483019190915260748201526094019392505050565b60208152600061152460208301846120ee565b60a08152600061220f60a08301886120ee565b6001600160a01b039690961660208301525060408101939093526060830191909152608090910152919050565b60a08152600061224f60a08301886120ee565b6001600160a01b039687166020840152949095166040820152606081019290925260809091015292915050565b60c08152600061228f60c08301896120ee565b6001600160a01b039788166020840152604083019690965250929094166060830152608082015260a00191909152919050565b6040815260006122d560408301856120ee565b905082151560208301529392505050565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b604051601f8201601f1916810167ffffffffffffffff81118282101715612346576123466124ee565b604052919050565b600067ffffffffffffffff821115612368576123686124ee565b5060051b60200190565b60008219821115612385576123856124ac565b500190565b600060ff821660ff84168060ff038211156123a7576123a76124ac565b019392505050565b6000826123be576123be6124c2565b500490565b60008160001904831182151516156123dd576123dd6124ac565b500290565b6000828210156123f4576123f46124ac565b500390565b600060ff821660ff841680821015612413576124136124ac565b90039392505050565b60005b8381101561243757818101518382015260200161241f565b838111156114b55750506000910152565b600181811c9082168061245c57607f821691505b60208210811415610df357634e487b7160e01b600052602260045260246000fd5b6000600019821415612491576124916124ac565b5060010190565b6000826124a7576124a76124c2565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461251957600080fd5b5056fe307830303030303030303030303030303030303030303030303030303030303030303030303030303030a2646970667358221220b0609db3eabb0f1f8782827f2bccb06d505bbc0d985a0f9d1bfefc33cea4880664736f6c63430008070033",
}

// TFuelTokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use TFuelTokenBankMetaData.ABI instead.
var TFuelTokenBankABI = TFuelTokenBankMetaData.ABI

// TFuelTokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TFuelTokenBankMetaData.Bin instead.
var TFuelTokenBankBin = TFuelTokenBankMetaData.Bin

// DeployTFuelTokenBank deploys a new Ethereum contract, binding an instance of TFuelTokenBank to it.
func DeployTFuelTokenBank(auth *bind.TransactOpts, backend bind.ContractBackend, mainchainid_ *big.Int, subchainregistrar_ common.Address) (common.Address, *types.Transaction, *TFuelTokenBank, error) {
	parsed, err := TFuelTokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TFuelTokenBankBin), backend, mainchainid_, subchainregistrar_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TFuelTokenBank{TFuelTokenBankCaller: TFuelTokenBankCaller{contract: contract}, TFuelTokenBankTransactor: TFuelTokenBankTransactor{contract: contract}, TFuelTokenBankFilterer: TFuelTokenBankFilterer{contract: contract}}, nil
}

// TFuelTokenBank is an auto generated Go binding around an Ethereum contract.
type TFuelTokenBank struct {
	TFuelTokenBankCaller     // Read-only binding to the contract
	TFuelTokenBankTransactor // Write-only binding to the contract
	TFuelTokenBankFilterer   // Log filterer for contract events
}

// TFuelTokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type TFuelTokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TFuelTokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TFuelTokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TFuelTokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TFuelTokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TFuelTokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TFuelTokenBankSession struct {
	Contract     *TFuelTokenBank   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TFuelTokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TFuelTokenBankCallerSession struct {
	Contract *TFuelTokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TFuelTokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TFuelTokenBankTransactorSession struct {
	Contract     *TFuelTokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TFuelTokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type TFuelTokenBankRaw struct {
	Contract *TFuelTokenBank // Generic contract binding to access the raw methods on
}

// TFuelTokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TFuelTokenBankCallerRaw struct {
	Contract *TFuelTokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// TFuelTokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TFuelTokenBankTransactorRaw struct {
	Contract *TFuelTokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTFuelTokenBank creates a new instance of TFuelTokenBank, bound to a specific deployed contract.
func NewTFuelTokenBank(address common.Address, backend bind.ContractBackend) (*TFuelTokenBank, error) {
	contract, err := bindTFuelTokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBank{TFuelTokenBankCaller: TFuelTokenBankCaller{contract: contract}, TFuelTokenBankTransactor: TFuelTokenBankTransactor{contract: contract}, TFuelTokenBankFilterer: TFuelTokenBankFilterer{contract: contract}}, nil
}

// NewTFuelTokenBankCaller creates a new read-only instance of TFuelTokenBank, bound to a specific deployed contract.
func NewTFuelTokenBankCaller(address common.Address, caller bind.ContractCaller) (*TFuelTokenBankCaller, error) {
	contract, err := bindTFuelTokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankCaller{contract: contract}, nil
}

// NewTFuelTokenBankTransactor creates a new write-only instance of TFuelTokenBank, bound to a specific deployed contract.
func NewTFuelTokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*TFuelTokenBankTransactor, error) {
	contract, err := bindTFuelTokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTransactor{contract: contract}, nil
}

// NewTFuelTokenBankFilterer creates a new log filterer instance of TFuelTokenBank, bound to a specific deployed contract.
func NewTFuelTokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*TFuelTokenBankFilterer, error) {
	contract, err := bindTFuelTokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankFilterer{contract: contract}, nil
}

// bindTFuelTokenBank binds a generic wrapper to an already deployed contract.
func bindTFuelTokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TFuelTokenBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TFuelTokenBank *TFuelTokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TFuelTokenBank.Contract.TFuelTokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TFuelTokenBank *TFuelTokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.TFuelTokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TFuelTokenBank *TFuelTokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.TFuelTokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TFuelTokenBank *TFuelTokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TFuelTokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TFuelTokenBank *TFuelTokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TFuelTokenBank *TFuelTokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.contract.Transact(opts, method, params...)
}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankCaller) Alldenoms(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "alldenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankSession) Alldenoms(arg0 *big.Int) (string, error) {
	return _TFuelTokenBank.Contract.Alldenoms(&_TFuelTokenBank.CallOpts, arg0)
}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Alldenoms(arg0 *big.Int) (string, error) {
	return _TFuelTokenBank.Contract.Alldenoms(&_TFuelTokenBank.CallOpts, arg0)
}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankCaller) Allvouchers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "allvouchers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankSession) Allvouchers(arg0 *big.Int) (common.Address, error) {
	return _TFuelTokenBank.Contract.Allvouchers(&_TFuelTokenBank.CallOpts, arg0)
}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Allvouchers(arg0 *big.Int) (common.Address, error) {
	return _TFuelTokenBank.Contract.Allvouchers(&_TFuelTokenBank.CallOpts, arg0)
}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_TFuelTokenBank *TFuelTokenBankCaller) Denomtovoucherlookup(opts *bind.CallOpts, arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "denomtovoucherlookup", arg0)

	outstruct := new(struct {
		Contractaddress common.Address
		Exists          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Contractaddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_TFuelTokenBank *TFuelTokenBankSession) Denomtovoucherlookup(arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	return _TFuelTokenBank.Contract.Denomtovoucherlookup(&_TFuelTokenBank.CallOpts, arg0)
}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Denomtovoucherlookup(arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	return _TFuelTokenBank.Contract.Denomtovoucherlookup(&_TFuelTokenBank.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCaller) Exists(opts *bind.CallOpts, denom string) (bool, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "exists", denom)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankSession) Exists(denom string) (bool, error) {
	return _TFuelTokenBank.Contract.Exists(&_TFuelTokenBank.CallOpts, denom)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Exists(denom string) (bool, error) {
	return _TFuelTokenBank.Contract.Exists(&_TFuelTokenBank.CallOpts, denom)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCaller) Exists0(opts *bind.CallOpts, voucheraddress common.Address) (bool, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "exists0", voucheraddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankSession) Exists0(voucheraddress common.Address) (bool, error) {
	return _TFuelTokenBank.Contract.Exists0(&_TFuelTokenBank.CallOpts, voucheraddress)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Exists0(voucheraddress common.Address) (bool, error) {
	return _TFuelTokenBank.Contract.Exists0(&_TFuelTokenBank.CallOpts, voucheraddress)
}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankCaller) Getdenom(opts *bind.CallOpts, vouchercontractaddr common.Address) (string, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "getdenom", vouchercontractaddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankSession) Getdenom(vouchercontractaddr common.Address) (string, error) {
	return _TFuelTokenBank.Contract.Getdenom(&_TFuelTokenBank.CallOpts, vouchercontractaddr)
}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Getdenom(vouchercontractaddr common.Address) (string, error) {
	return _TFuelTokenBank.Contract.Getdenom(&_TFuelTokenBank.CallOpts, vouchercontractaddr)
}

// Getmaxprocessedtokenlocknonce is a free data retrieval call binding the contract method 0x46e241e5.
//
// Solidity: function getmaxprocessedtokenlocknonce(uint256 chainid) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Getmaxprocessedtokenlocknonce(opts *bind.CallOpts, chainid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "getmaxprocessedtokenlocknonce", chainid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getmaxprocessedtokenlocknonce is a free data retrieval call binding the contract method 0x46e241e5.
//
// Solidity: function getmaxprocessedtokenlocknonce(uint256 chainid) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Getmaxprocessedtokenlocknonce(chainid *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Getmaxprocessedtokenlocknonce(&_TFuelTokenBank.CallOpts, chainid)
}

// Getmaxprocessedtokenlocknonce is a free data retrieval call binding the contract method 0x46e241e5.
//
// Solidity: function getmaxprocessedtokenlocknonce(uint256 chainid) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Getmaxprocessedtokenlocknonce(chainid *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Getmaxprocessedtokenlocknonce(&_TFuelTokenBank.CallOpts, chainid)
}

// Getmaxprocessedvoucherburnnonce is a free data retrieval call binding the contract method 0xd0f2c8f0.
//
// Solidity: function getmaxprocessedvoucherburnnonce(uint256 chainid) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Getmaxprocessedvoucherburnnonce(opts *bind.CallOpts, chainid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "getmaxprocessedvoucherburnnonce", chainid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getmaxprocessedvoucherburnnonce is a free data retrieval call binding the contract method 0xd0f2c8f0.
//
// Solidity: function getmaxprocessedvoucherburnnonce(uint256 chainid) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Getmaxprocessedvoucherburnnonce(chainid *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Getmaxprocessedvoucherburnnonce(&_TFuelTokenBank.CallOpts, chainid)
}

// Getmaxprocessedvoucherburnnonce is a free data retrieval call binding the contract method 0xd0f2c8f0.
//
// Solidity: function getmaxprocessedvoucherburnnonce(uint256 chainid) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Getmaxprocessedvoucherburnnonce(chainid *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Getmaxprocessedvoucherburnnonce(&_TFuelTokenBank.CallOpts, chainid)
}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankCaller) Getvoucher(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "getvoucher", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankSession) Getvoucher(denom string) (common.Address, error) {
	return _TFuelTokenBank.Contract.Getvoucher(&_TFuelTokenBank.CallOpts, denom)
}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Getvoucher(denom string) (common.Address, error) {
	return _TFuelTokenBank.Contract.Getvoucher(&_TFuelTokenBank.CallOpts, denom)
}

// Isonmainchain is a free data retrieval call binding the contract method 0x185eb163.
//
// Solidity: function isonmainchain() view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCaller) Isonmainchain(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "isonmainchain")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Isonmainchain is a free data retrieval call binding the contract method 0x185eb163.
//
// Solidity: function isonmainchain() view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankSession) Isonmainchain() (bool, error) {
	return _TFuelTokenBank.Contract.Isonmainchain(&_TFuelTokenBank.CallOpts)
}

// Isonmainchain is a free data retrieval call binding the contract method 0x185eb163.
//
// Solidity: function isonmainchain() view returns(bool)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Isonmainchain() (bool, error) {
	return _TFuelTokenBank.Contract.Isonmainchain(&_TFuelTokenBank.CallOpts)
}

// Mainchainid is a free data retrieval call binding the contract method 0xba3599b3.
//
// Solidity: function mainchainid() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Mainchainid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "mainchainid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mainchainid is a free data retrieval call binding the contract method 0xba3599b3.
//
// Solidity: function mainchainid() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Mainchainid() (*big.Int, error) {
	return _TFuelTokenBank.Contract.Mainchainid(&_TFuelTokenBank.CallOpts)
}

// Mainchainid is a free data retrieval call binding the contract method 0xba3599b3.
//
// Solidity: function mainchainid() view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Mainchainid() (*big.Int, error) {
	return _TFuelTokenBank.Contract.Mainchainid(&_TFuelTokenBank.CallOpts)
}

// Tokenlocknoncemap is a free data retrieval call binding the contract method 0x52baad62.
//
// Solidity: function tokenlocknoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Tokenlocknoncemap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "tokenlocknoncemap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokenlocknoncemap is a free data retrieval call binding the contract method 0x52baad62.
//
// Solidity: function tokenlocknoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Tokenlocknoncemap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Tokenlocknoncemap(&_TFuelTokenBank.CallOpts, arg0)
}

// Tokenlocknoncemap is a free data retrieval call binding the contract method 0x52baad62.
//
// Solidity: function tokenlocknoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Tokenlocknoncemap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Tokenlocknoncemap(&_TFuelTokenBank.CallOpts, arg0)
}

// Tokenlockvotingrecords is a free data retrieval call binding the contract method 0x1b637e32.
//
// Solidity: function tokenlockvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_TFuelTokenBank *TFuelTokenBankCaller) Tokenlockvotingrecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "tokenlockvotingrecords", arg0, arg1)

	outstruct := new(struct {
		Dynasty          *big.Int
		Accumlatedshares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dynasty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Accumlatedshares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tokenlockvotingrecords is a free data retrieval call binding the contract method 0x1b637e32.
//
// Solidity: function tokenlockvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_TFuelTokenBank *TFuelTokenBankSession) Tokenlockvotingrecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	return _TFuelTokenBank.Contract.Tokenlockvotingrecords(&_TFuelTokenBank.CallOpts, arg0, arg1)
}

// Tokenlockvotingrecords is a free data retrieval call binding the contract method 0x1b637e32.
//
// Solidity: function tokenlockvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Tokenlockvotingrecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	return _TFuelTokenBank.Contract.Tokenlockvotingrecords(&_TFuelTokenBank.CallOpts, arg0, arg1)
}

// Tokenunlocknoncemap is a free data retrieval call binding the contract method 0x12d53e08.
//
// Solidity: function tokenunlocknoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Tokenunlocknoncemap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "tokenunlocknoncemap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokenunlocknoncemap is a free data retrieval call binding the contract method 0x12d53e08.
//
// Solidity: function tokenunlocknoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Tokenunlocknoncemap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Tokenunlocknoncemap(&_TFuelTokenBank.CallOpts, arg0)
}

// Tokenunlocknoncemap is a free data retrieval call binding the contract method 0x12d53e08.
//
// Solidity: function tokenunlocknoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Tokenunlocknoncemap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Tokenunlocknoncemap(&_TFuelTokenBank.CallOpts, arg0)
}

// Totallockedamounts is a free data retrieval call binding the contract method 0x5a35e496.
//
// Solidity: function totallockedamounts(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Totallockedamounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "totallockedamounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Totallockedamounts is a free data retrieval call binding the contract method 0x5a35e496.
//
// Solidity: function totallockedamounts(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Totallockedamounts(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Totallockedamounts(&_TFuelTokenBank.CallOpts, arg0)
}

// Totallockedamounts is a free data retrieval call binding the contract method 0x5a35e496.
//
// Solidity: function totallockedamounts(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Totallockedamounts(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Totallockedamounts(&_TFuelTokenBank.CallOpts, arg0)
}

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_TFuelTokenBank *TFuelTokenBankCaller) Voucheraddresstodenomlookup(opts *bind.CallOpts, arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "voucheraddresstodenomlookup", arg0)

	outstruct := new(struct {
		Denom  string
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Denom = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_TFuelTokenBank *TFuelTokenBankSession) Voucheraddresstodenomlookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _TFuelTokenBank.Contract.Voucheraddresstodenomlookup(&_TFuelTokenBank.CallOpts, arg0)
}

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Voucheraddresstodenomlookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _TFuelTokenBank.Contract.Voucheraddresstodenomlookup(&_TFuelTokenBank.CallOpts, arg0)
}

// Voucherburnnoncemap is a free data retrieval call binding the contract method 0x617d9140.
//
// Solidity: function voucherburnnoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Voucherburnnoncemap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "voucherburnnoncemap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Voucherburnnoncemap is a free data retrieval call binding the contract method 0x617d9140.
//
// Solidity: function voucherburnnoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Voucherburnnoncemap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Voucherburnnoncemap(&_TFuelTokenBank.CallOpts, arg0)
}

// Voucherburnnoncemap is a free data retrieval call binding the contract method 0x617d9140.
//
// Solidity: function voucherburnnoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Voucherburnnoncemap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Voucherburnnoncemap(&_TFuelTokenBank.CallOpts, arg0)
}

// Voucherburnvotingrecords is a free data retrieval call binding the contract method 0x147ea516.
//
// Solidity: function voucherburnvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_TFuelTokenBank *TFuelTokenBankCaller) Voucherburnvotingrecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "voucherburnvotingrecords", arg0, arg1)

	outstruct := new(struct {
		Dynasty          *big.Int
		Accumlatedshares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dynasty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Accumlatedshares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voucherburnvotingrecords is a free data retrieval call binding the contract method 0x147ea516.
//
// Solidity: function voucherburnvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_TFuelTokenBank *TFuelTokenBankSession) Voucherburnvotingrecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	return _TFuelTokenBank.Contract.Voucherburnvotingrecords(&_TFuelTokenBank.CallOpts, arg0, arg1)
}

// Voucherburnvotingrecords is a free data retrieval call binding the contract method 0x147ea516.
//
// Solidity: function voucherburnvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Voucherburnvotingrecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	return _TFuelTokenBank.Contract.Voucherburnvotingrecords(&_TFuelTokenBank.CallOpts, arg0, arg1)
}

// Vouchermintnoncemap is a free data retrieval call binding the contract method 0xe7614834.
//
// Solidity: function vouchermintnoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCaller) Vouchermintnoncemap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TFuelTokenBank.contract.Call(opts, &out, "vouchermintnoncemap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vouchermintnoncemap is a free data retrieval call binding the contract method 0xe7614834.
//
// Solidity: function vouchermintnoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankSession) Vouchermintnoncemap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Vouchermintnoncemap(&_TFuelTokenBank.CallOpts, arg0)
}

// Vouchermintnoncemap is a free data retrieval call binding the contract method 0xe7614834.
//
// Solidity: function vouchermintnoncemap(uint256 ) view returns(uint256)
func (_TFuelTokenBank *TFuelTokenBankCallerSession) Vouchermintnoncemap(arg0 *big.Int) (*big.Int, error) {
	return _TFuelTokenBank.Contract.Vouchermintnoncemap(&_TFuelTokenBank.CallOpts, arg0)
}

// Burnvouchers is a paid mutator transaction binding the contract method 0xc18b2085.
//
// Solidity: function burnvouchers(address targetchaintokenreceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactor) Burnvouchers(opts *bind.TransactOpts, targetchaintokenreceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.contract.Transact(opts, "burnvouchers", targetchaintokenreceiver)
}

// Burnvouchers is a paid mutator transaction binding the contract method 0xc18b2085.
//
// Solidity: function burnvouchers(address targetchaintokenreceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankSession) Burnvouchers(targetchaintokenreceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.Burnvouchers(&_TFuelTokenBank.TransactOpts, targetchaintokenreceiver)
}

// Burnvouchers is a paid mutator transaction binding the contract method 0xc18b2085.
//
// Solidity: function burnvouchers(address targetchaintokenreceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactorSession) Burnvouchers(targetchaintokenreceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.Burnvouchers(&_TFuelTokenBank.TransactOpts, targetchaintokenreceiver)
}

// Locktokens is a paid mutator transaction binding the contract method 0xd9a3958f.
//
// Solidity: function locktokens(uint256 targetchainid, address targetchainvoucherreceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactor) Locktokens(opts *bind.TransactOpts, targetchainid *big.Int, targetchainvoucherreceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.contract.Transact(opts, "locktokens", targetchainid, targetchainvoucherreceiver)
}

// Locktokens is a paid mutator transaction binding the contract method 0xd9a3958f.
//
// Solidity: function locktokens(uint256 targetchainid, address targetchainvoucherreceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankSession) Locktokens(targetchainid *big.Int, targetchainvoucherreceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.Locktokens(&_TFuelTokenBank.TransactOpts, targetchainid, targetchainvoucherreceiver)
}

// Locktokens is a paid mutator transaction binding the contract method 0xd9a3958f.
//
// Solidity: function locktokens(uint256 targetchainid, address targetchainvoucherreceiver) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactorSession) Locktokens(targetchainid *big.Int, targetchainvoucherreceiver common.Address) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.Locktokens(&_TFuelTokenBank.TransactOpts, targetchainid, targetchainvoucherreceiver)
}

// Mintvouchers is a paid mutator transaction binding the contract method 0x857496f6.
//
// Solidity: function mintvouchers(string denom, address targetchainvoucherreceiver, uint256 mintedamount, uint256 dynasty, uint256 sourcechaintokenlocknonce) returns()
func (_TFuelTokenBank *TFuelTokenBankTransactor) Mintvouchers(opts *bind.TransactOpts, denom string, targetchainvoucherreceiver common.Address, mintedamount *big.Int, dynasty *big.Int, sourcechaintokenlocknonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.contract.Transact(opts, "mintvouchers", denom, targetchainvoucherreceiver, mintedamount, dynasty, sourcechaintokenlocknonce)
}

// Mintvouchers is a paid mutator transaction binding the contract method 0x857496f6.
//
// Solidity: function mintvouchers(string denom, address targetchainvoucherreceiver, uint256 mintedamount, uint256 dynasty, uint256 sourcechaintokenlocknonce) returns()
func (_TFuelTokenBank *TFuelTokenBankSession) Mintvouchers(denom string, targetchainvoucherreceiver common.Address, mintedamount *big.Int, dynasty *big.Int, sourcechaintokenlocknonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.Mintvouchers(&_TFuelTokenBank.TransactOpts, denom, targetchainvoucherreceiver, mintedamount, dynasty, sourcechaintokenlocknonce)
}

// Mintvouchers is a paid mutator transaction binding the contract method 0x857496f6.
//
// Solidity: function mintvouchers(string denom, address targetchainvoucherreceiver, uint256 mintedamount, uint256 dynasty, uint256 sourcechaintokenlocknonce) returns()
func (_TFuelTokenBank *TFuelTokenBankTransactorSession) Mintvouchers(denom string, targetchainvoucherreceiver common.Address, mintedamount *big.Int, dynasty *big.Int, sourcechaintokenlocknonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.Mintvouchers(&_TFuelTokenBank.TransactOpts, denom, targetchainvoucherreceiver, mintedamount, dynasty, sourcechaintokenlocknonce)
}

// Unlocktokens is a paid mutator transaction binding the contract method 0x9fc3fc8a.
//
// Solidity: function unlocktokens(uint256 sourcechainid, address targetchaintokenreceiver, uint256 unlockamount, uint256 dynasty, uint256 sourcechainvoucherburnnonce) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactor) Unlocktokens(opts *bind.TransactOpts, sourcechainid *big.Int, targetchaintokenreceiver common.Address, unlockamount *big.Int, dynasty *big.Int, sourcechainvoucherburnnonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.contract.Transact(opts, "unlocktokens", sourcechainid, targetchaintokenreceiver, unlockamount, dynasty, sourcechainvoucherburnnonce)
}

// Unlocktokens is a paid mutator transaction binding the contract method 0x9fc3fc8a.
//
// Solidity: function unlocktokens(uint256 sourcechainid, address targetchaintokenreceiver, uint256 unlockamount, uint256 dynasty, uint256 sourcechainvoucherburnnonce) payable returns()
func (_TFuelTokenBank *TFuelTokenBankSession) Unlocktokens(sourcechainid *big.Int, targetchaintokenreceiver common.Address, unlockamount *big.Int, dynasty *big.Int, sourcechainvoucherburnnonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.Unlocktokens(&_TFuelTokenBank.TransactOpts, sourcechainid, targetchaintokenreceiver, unlockamount, dynasty, sourcechainvoucherburnnonce)
}

// Unlocktokens is a paid mutator transaction binding the contract method 0x9fc3fc8a.
//
// Solidity: function unlocktokens(uint256 sourcechainid, address targetchaintokenreceiver, uint256 unlockamount, uint256 dynasty, uint256 sourcechainvoucherburnnonce) payable returns()
func (_TFuelTokenBank *TFuelTokenBankTransactorSession) Unlocktokens(sourcechainid *big.Int, targetchaintokenreceiver common.Address, unlockamount *big.Int, dynasty *big.Int, sourcechainvoucherburnnonce *big.Int) (*types.Transaction, error) {
	return _TFuelTokenBank.Contract.Unlocktokens(&_TFuelTokenBank.TransactOpts, sourcechainid, targetchaintokenreceiver, unlockamount, dynasty, sourcechainvoucherburnnonce)
}

// TFuelTokenBankTfueltokenlockedIterator is returned from FilterTfueltokenlocked and is used to iterate over the raw logs and unpacked data for Tfueltokenlocked events raised by the TFuelTokenBank contract.
type TFuelTokenBankTfueltokenlockedIterator struct {
	Event *TFuelTokenBankTfueltokenlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TFuelTokenBankTfueltokenlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankTfueltokenlocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TFuelTokenBankTfueltokenlocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TFuelTokenBankTfueltokenlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankTfueltokenlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankTfueltokenlocked represents a Tfueltokenlocked event raised by the TFuelTokenBank contract.
type TFuelTokenBankTfueltokenlocked struct {
	Denom                      string
	Sourcechaintokensender     common.Address
	Targetchainid              *big.Int
	Targetchainvoucherreceiver common.Address
	Lockedamount               *big.Int
	Tokenlocknonce             *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTfueltokenlocked is a free log retrieval operation binding the contract event 0xdb6db927072dc34c88f44b1df7298a595be64983ad64eb1e5506f13bb1c72696.
//
// Solidity: event tfueltokenlocked(string denom, address sourcechaintokensender, uint256 targetchainid, address targetchainvoucherreceiver, uint256 lockedamount, uint256 tokenlocknonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterTfueltokenlocked(opts *bind.FilterOpts) (*TFuelTokenBankTfueltokenlockedIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "tfueltokenlocked")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTfueltokenlockedIterator{contract: _TFuelTokenBank.contract, event: "tfueltokenlocked", logs: logs, sub: sub}, nil
}

// WatchTfueltokenlocked is a free log subscription operation binding the contract event 0xdb6db927072dc34c88f44b1df7298a595be64983ad64eb1e5506f13bb1c72696.
//
// Solidity: event tfueltokenlocked(string denom, address sourcechaintokensender, uint256 targetchainid, address targetchainvoucherreceiver, uint256 lockedamount, uint256 tokenlocknonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchTfueltokenlocked(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankTfueltokenlocked) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "tfueltokenlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankTfueltokenlocked)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "tfueltokenlocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTfueltokenlocked is a log parse operation binding the contract event 0xdb6db927072dc34c88f44b1df7298a595be64983ad64eb1e5506f13bb1c72696.
//
// Solidity: event tfueltokenlocked(string denom, address sourcechaintokensender, uint256 targetchainid, address targetchainvoucherreceiver, uint256 lockedamount, uint256 tokenlocknonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseTfueltokenlocked(log types.Log) (*TFuelTokenBankTfueltokenlocked, error) {
	event := new(TFuelTokenBankTfueltokenlocked)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "tfueltokenlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TFuelTokenBankTfueltokenunlockedIterator is returned from FilterTfueltokenunlocked and is used to iterate over the raw logs and unpacked data for Tfueltokenunlocked events raised by the TFuelTokenBank contract.
type TFuelTokenBankTfueltokenunlockedIterator struct {
	Event *TFuelTokenBankTfueltokenunlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TFuelTokenBankTfueltokenunlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankTfueltokenunlocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TFuelTokenBankTfueltokenunlocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TFuelTokenBankTfueltokenunlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankTfueltokenunlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankTfueltokenunlocked represents a Tfueltokenunlocked event raised by the TFuelTokenBank contract.
type TFuelTokenBankTfueltokenunlocked struct {
	Denom                       string
	Targetchaintokenreceiver    common.Address
	Unlockedamount              *big.Int
	Sourcechainvoucherburnnonce *big.Int
	Tokenunlocknonce            *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTfueltokenunlocked is a free log retrieval operation binding the contract event 0x3142136ad28970fc04f40fb3af4e056d396a71084258e0b5897ff4f42a4609dc.
//
// Solidity: event tfueltokenunlocked(string denom, address targetchaintokenreceiver, uint256 unlockedamount, uint256 sourcechainvoucherburnnonce, uint256 tokenunlocknonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterTfueltokenunlocked(opts *bind.FilterOpts) (*TFuelTokenBankTfueltokenunlockedIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "tfueltokenunlocked")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTfueltokenunlockedIterator{contract: _TFuelTokenBank.contract, event: "tfueltokenunlocked", logs: logs, sub: sub}, nil
}

// WatchTfueltokenunlocked is a free log subscription operation binding the contract event 0x3142136ad28970fc04f40fb3af4e056d396a71084258e0b5897ff4f42a4609dc.
//
// Solidity: event tfueltokenunlocked(string denom, address targetchaintokenreceiver, uint256 unlockedamount, uint256 sourcechainvoucherburnnonce, uint256 tokenunlocknonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchTfueltokenunlocked(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankTfueltokenunlocked) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "tfueltokenunlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankTfueltokenunlocked)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "tfueltokenunlocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTfueltokenunlocked is a log parse operation binding the contract event 0x3142136ad28970fc04f40fb3af4e056d396a71084258e0b5897ff4f42a4609dc.
//
// Solidity: event tfueltokenunlocked(string denom, address targetchaintokenreceiver, uint256 unlockedamount, uint256 sourcechainvoucherburnnonce, uint256 tokenunlocknonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseTfueltokenunlocked(log types.Log) (*TFuelTokenBankTfueltokenunlocked, error) {
	event := new(TFuelTokenBankTfueltokenunlocked)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "tfueltokenunlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TFuelTokenBankTfuelvoucherburnedIterator is returned from FilterTfuelvoucherburned and is used to iterate over the raw logs and unpacked data for Tfuelvoucherburned events raised by the TFuelTokenBank contract.
type TFuelTokenBankTfuelvoucherburnedIterator struct {
	Event *TFuelTokenBankTfuelvoucherburned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TFuelTokenBankTfuelvoucherburnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankTfuelvoucherburned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TFuelTokenBankTfuelvoucherburned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TFuelTokenBankTfuelvoucherburnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankTfuelvoucherburnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankTfuelvoucherburned represents a Tfuelvoucherburned event raised by the TFuelTokenBank contract.
type TFuelTokenBankTfuelvoucherburned struct {
	Denom                    string
	Sourcechainvoucherowner  common.Address
	Targetchaintokenreceiver common.Address
	Burnedamount             *big.Int
	Voucherburnnonce         *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterTfuelvoucherburned is a free log retrieval operation binding the contract event 0x543e645a5be2a24be930fbb96caa6d30176fc20d2609d48cbbef0ed5d9ad3ac5.
//
// Solidity: event tfuelvoucherburned(string denom, address sourcechainvoucherowner, address targetchaintokenreceiver, uint256 burnedamount, uint256 voucherburnnonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterTfuelvoucherburned(opts *bind.FilterOpts) (*TFuelTokenBankTfuelvoucherburnedIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "tfuelvoucherburned")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTfuelvoucherburnedIterator{contract: _TFuelTokenBank.contract, event: "tfuelvoucherburned", logs: logs, sub: sub}, nil
}

// WatchTfuelvoucherburned is a free log subscription operation binding the contract event 0x543e645a5be2a24be930fbb96caa6d30176fc20d2609d48cbbef0ed5d9ad3ac5.
//
// Solidity: event tfuelvoucherburned(string denom, address sourcechainvoucherowner, address targetchaintokenreceiver, uint256 burnedamount, uint256 voucherburnnonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchTfuelvoucherburned(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankTfuelvoucherburned) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "tfuelvoucherburned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankTfuelvoucherburned)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "tfuelvoucherburned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTfuelvoucherburned is a log parse operation binding the contract event 0x543e645a5be2a24be930fbb96caa6d30176fc20d2609d48cbbef0ed5d9ad3ac5.
//
// Solidity: event tfuelvoucherburned(string denom, address sourcechainvoucherowner, address targetchaintokenreceiver, uint256 burnedamount, uint256 voucherburnnonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseTfuelvoucherburned(log types.Log) (*TFuelTokenBankTfuelvoucherburned, error) {
	event := new(TFuelTokenBankTfuelvoucherburned)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "tfuelvoucherburned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TFuelTokenBankTfuelvouchermintedIterator is returned from FilterTfuelvoucherminted and is used to iterate over the raw logs and unpacked data for Tfuelvoucherminted events raised by the TFuelTokenBank contract.
type TFuelTokenBankTfuelvouchermintedIterator struct {
	Event *TFuelTokenBankTfuelvoucherminted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TFuelTokenBankTfuelvouchermintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TFuelTokenBankTfuelvoucherminted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TFuelTokenBankTfuelvoucherminted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TFuelTokenBankTfuelvouchermintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TFuelTokenBankTfuelvouchermintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TFuelTokenBankTfuelvoucherminted represents a Tfuelvoucherminted event raised by the TFuelTokenBank contract.
type TFuelTokenBankTfuelvoucherminted struct {
	Denom                      string
	Targetchainvoucherreceiver common.Address
	Mintedamount               *big.Int
	Sourcechaintokenlocknonce  *big.Int
	Vouchermintnonce           *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTfuelvoucherminted is a free log retrieval operation binding the contract event 0x1e3ac0ceace51a288a5351c90c91ea8f20fd784bedd512ba699acc90cf97a356.
//
// Solidity: event tfuelvoucherminted(string denom, address targetchainvoucherreceiver, uint256 mintedamount, uint256 sourcechaintokenlocknonce, uint256 vouchermintnonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) FilterTfuelvoucherminted(opts *bind.FilterOpts) (*TFuelTokenBankTfuelvouchermintedIterator, error) {

	logs, sub, err := _TFuelTokenBank.contract.FilterLogs(opts, "tfuelvoucherminted")
	if err != nil {
		return nil, err
	}
	return &TFuelTokenBankTfuelvouchermintedIterator{contract: _TFuelTokenBank.contract, event: "tfuelvoucherminted", logs: logs, sub: sub}, nil
}

// WatchTfuelvoucherminted is a free log subscription operation binding the contract event 0x1e3ac0ceace51a288a5351c90c91ea8f20fd784bedd512ba699acc90cf97a356.
//
// Solidity: event tfuelvoucherminted(string denom, address targetchainvoucherreceiver, uint256 mintedamount, uint256 sourcechaintokenlocknonce, uint256 vouchermintnonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) WatchTfuelvoucherminted(opts *bind.WatchOpts, sink chan<- *TFuelTokenBankTfuelvoucherminted) (event.Subscription, error) {

	logs, sub, err := _TFuelTokenBank.contract.WatchLogs(opts, "tfuelvoucherminted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TFuelTokenBankTfuelvoucherminted)
				if err := _TFuelTokenBank.contract.UnpackLog(event, "tfuelvoucherminted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTfuelvoucherminted is a log parse operation binding the contract event 0x1e3ac0ceace51a288a5351c90c91ea8f20fd784bedd512ba699acc90cf97a356.
//
// Solidity: event tfuelvoucherminted(string denom, address targetchainvoucherreceiver, uint256 mintedamount, uint256 sourcechaintokenlocknonce, uint256 vouchermintnonce)
func (_TFuelTokenBank *TFuelTokenBankFilterer) ParseTfuelvoucherminted(log types.Log) (*TFuelTokenBankTfuelvoucherminted, error) {
	event := new(TFuelTokenBankTfuelvoucherminted)
	if err := _TFuelTokenBank.contract.UnpackLog(event, "tfuelvoucherminted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}