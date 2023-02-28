// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executor

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// ExecutorMetaData contains all meta data concerning the Executor contract.
var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"nft_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receipient\",\"type\":\"address\"}],\"name\":\"buyImmediate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyImmediate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receipient\",\"type\":\"address\"}],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"contributionAmt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200169b3803806200169b833981016040819052620000349162000163565b828260036200004483826200027f565b5060046200005382826200027f565b505060006007819055600a80546001600160a01b039094166001600160a01b03199485168117909155600b805490941617909255506008805460ff19169055600955506200034b9050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000c657600080fd5b81516001600160401b0380821115620000e357620000e36200009e565b604051601f8301601f19908116603f011681019082821181831017156200010e576200010e6200009e565b816040528381526020925086838588010111156200012b57600080fd5b600091505b838210156200014f578582018301518183018401529082019062000130565b600093810190920192909252949350505050565b6000806000606084860312156200017957600080fd5b83516001600160401b03808211156200019157600080fd5b6200019f87838801620000b4565b94506020860151915080821115620001b657600080fd5b50620001c586828701620000b4565b604086015190935090506001600160a01b0381168114620001e557600080fd5b809150509250925092565b600181811c908216806200020557607f821691505b6020821081036200022657634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200027a57600081815260208120601f850160051c81016020861015620002555750805b601f850160051c820191505b81811015620002765782815560010162000261565b5050505b505050565b81516001600160401b038111156200029b576200029b6200009e565b620002b381620002ac8454620001f0565b846200022c565b602080601f831160018114620002eb5760008415620002d25750858301515b600019600386901b1c1916600185901b17855562000276565b600085815260208120601f198616915b828110156200031c57888601518255948401946001909101908401620002fb565b50858210156200033b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611340806200035b6000396000f3fe60806040526004361061012a5760003560e01c806368428a1b116100ab57806395d89b411161006f57806395d89b4114610322578063a035b1fe14610337578063a457c2d71461034c578063a9059cbb1461036c578063d7bb99ba1461038c578063dd62ed3e1461039457600080fd5b806368428a1b146102865780636f24539f1461029b57806370a08231146102a357806373e888fd146102d957806376c6c2ef146102ec57600080fd5b806323b872dd116100f257806323b872dd146101f5578063278ecde114610215578063313ce5671461023757806339509351146102535780635b286a121461027357600080fd5b806306fdde031461012f578063095ea7b31461015a57806312065fe01461018a578063150b7a02146101a757806318160ddd146101e0575b600080fd5b34801561013b57600080fd5b506101446103b4565b6040516101519190611067565b60405180910390f35b34801561016657600080fd5b5061017a6101753660046110d1565b610446565b6040519015158152602001610151565b34801561019657600080fd5b50475b604051908152602001610151565b3480156101b357600080fd5b506101c76101c23660046110fb565b610460565b6040516001600160e01b03199091168152602001610151565b3480156101ec57600080fd5b50600254610199565b34801561020157600080fd5b5061017a610210366004611196565b6104ef565b34801561022157600080fd5b506102356102303660046111d2565b610513565b005b34801561024357600080fd5b5060405160128152602001610151565b34801561025f57600080fd5b5061017a61026e3660046110d1565b6106bd565b6102356102813660046111eb565b6106df565b34801561029257600080fd5b5061017a61084e565b6102356108c1565b3480156102af57600080fd5b506101996102be3660046111eb565b6001600160a01b031660009081526020819052604090205490565b6102356102e73660046111eb565b6108cc565b3480156102f857600080fd5b506101996103073660046111eb565b6001600160a01b031660009081526005602052604090205490565b34801561032e57600080fd5b50610144610b2c565b34801561034357600080fd5b50610199610b3b565b34801561035857600080fd5b5061017a6103673660046110d1565b610ba9565b34801561037857600080fd5b5061017a6103873660046110d1565b610c24565b610235610c32565b3480156103a057600080fd5b506101996103af36600461120d565b610c3b565b6060600380546103c390611240565b80601f01602080910402602001604051908101604052809291908181526020018280546103ef90611240565b801561043c5780601f106104115761010080835404028352916020019161043c565b820191906000526020600020905b81548152906001019060200180831161041f57829003601f168201915b5050505050905090565b600033610454818585610c66565b60019150505b92915050565b600b546000906001600160a01b038781169116146104dd5760405162461bcd60e51b815260206004820152602f60248201527f43616e6e6f742072656365697665204552433732312066726f6d20756e72656760448201526e69737465726564206164647265737360881b60648201526084015b60405180910390fd5b50630a85bd0160e11b95945050505050565b6000336104fd858285610d8a565b610508858585610e04565b506001949350505050565b61051b61084e565b80610538575061052961084e565b158015610538575060085460ff165b6105845760405162461bcd60e51b815260206004820152601c60248201527f43616e6e6f7420726566756e642066696e616c697a65642073616c650000000060448201526064016104d4565b600081116105e05760405162461bcd60e51b8152602060048201526024808201527f526566756e6420616d6f756e74206d757374206265206d6f7265207468616e206044820152637a65726f60e01b60648201526084016104d4565b3360009081526005602052604090205481111561064f5760405162461bcd60e51b815260206004820152602760248201527f526566756e642063616e6e6f742065786365656420636f6e747269627574656460448201526608185b5bdd5b9d60ca1b60648201526084016104d4565b604051339082156108fc029083906000818181858888f1935050505015801561067c573d6000803e3d6000fd5b50336000908152600560205260408120805483929061069c908490611290565b9250508190555080600960008282546106b59190611290565b909155505050565b6000336104548185856106d08383610c3b565b6106da91906112a3565b610c66565b6106e761084e565b6107285760405162461bcd60e51b815260206004820152601260248201527153616c65206973206e6f742061637469766560701b60448201526064016104d4565b610730610b3b565b3410156107985760405162461bcd60e51b815260206004820152603060248201527f4d73672076616c7565206d7573742062652067726561746572207468616e206f60448201526f7220657175616c20746f20707269636560801b60648201526084016104d4565b600a5460405163f088d54760e01b81526001600160a01b0383811660048301529091169063f088d54790349060240160206040518083038185885af11580156107e5573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061080a91906112b6565b61083e5760405133903480156108fc02916000818181858888f1935050505015801561083a573d6000803e3d6000fd5b5050565b506008805460ff19166001179055565b600a54604080516368428a1b60e01b815290516000926001600160a01b0316916368428a1b9160048083019260209291908290030181865afa158015610898573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bc91906112b6565b905090565b6108ca336106df565b565b6108d461084e565b6109155760405162461bcd60e51b815260206004820152601260248201527153616c65206973206e6f742061637469766560701b60448201526064016104d4565b600034116109715760405162461bcd60e51b815260206004820152602360248201527f436f6e747269627574696f6e206d75737420626520677265617465722074686160448201526206e20360ec1b60648201526084016104d4565b6000805b6007548110156109bb576000818152600660205260409020546001600160a01b038085169116036109a957600191506109bb565b806109b3816112d8565b915050610975565b5080610a0157600780549060006109d1836112d8565b9091555050600754600090815260066020526040902080546001600160a01b0319166001600160a01b0384161790555b6001600160a01b03821660009081526005602052604081208054349290610a299084906112a3565b925050819055503460096000828254610a4291906112a3565b90915550610a509050610b3b565b6009541015610a5d575050565b600a5460095460405163f088d54760e01b81523060048201526001600160a01b039092169163f088d547919060240160206040518083038185885af1158015610aaa573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610acf91906112b6565b610ad7575050565b60005b600754811015610b27576000818152600660209081526040808320546001600160a01b0316808452600590925290912054610b159190610fa8565b80610b1f816112d8565b915050610ada565b505050565b6060600480546103c390611240565b600a546040805163501ad8ff60e11b815290516000926001600160a01b03169163a035b1fe9160048083019260209291908290030181865afa158015610b85573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bc91906112f1565b60003381610bb78286610c3b565b905083811015610c175760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016104d4565b6105088286868403610c66565b600033610454818585610e04565b6108ca336108cc565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b038316610cc85760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016104d4565b6001600160a01b038216610d295760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016104d4565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6000610d968484610c3b565b90506000198114610dfe5781811015610df15760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016104d4565b610dfe8484848403610c66565b50505050565b6001600160a01b038316610e685760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016104d4565b6001600160a01b038216610eca5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016104d4565b6001600160a01b03831660009081526020819052604090205481811015610f425760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016104d4565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610dfe565b6001600160a01b038216610ffe5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016104d4565b806002600082825461101091906112a3565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600060208083528351808285015260005b8181101561109457858101830151858201604001528201611078565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146110cc57600080fd5b919050565b600080604083850312156110e457600080fd5b6110ed836110b5565b946020939093013593505050565b60008060008060006080868803121561111357600080fd5b61111c866110b5565b945061112a602087016110b5565b935060408601359250606086013567ffffffffffffffff8082111561114e57600080fd5b818801915088601f83011261116257600080fd5b81358181111561117157600080fd5b89602082850101111561118357600080fd5b9699959850939650602001949392505050565b6000806000606084860312156111ab57600080fd5b6111b4846110b5565b92506111c2602085016110b5565b9150604084013590509250925092565b6000602082840312156111e457600080fd5b5035919050565b6000602082840312156111fd57600080fd5b611206826110b5565b9392505050565b6000806040838503121561122057600080fd5b611229836110b5565b9150611237602084016110b5565b90509250929050565b600181811c9082168061125457607f821691505b60208210810361127457634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561045a5761045a61127a565b8082018082111561045a5761045a61127a565b6000602082840312156112c857600080fd5b8151801515811461120657600080fd5b6000600182016112ea576112ea61127a565b5060010190565b60006020828403121561130357600080fd5b505191905056fea26469706673582212206b5af1b02cf5e960ff33a99cc98cd0b7c59cc5322f3145805cd0980e058ad30764736f6c63430008120033",
}

// ExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutorMetaData.ABI instead.
var ExecutorABI = ExecutorMetaData.ABI

// ExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutorMetaData.Bin instead.
var ExecutorBin = ExecutorMetaData.Bin

// DeployExecutor deploys a new Ethereum contract, binding an instance of Executor to it.
func DeployExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, nft_ common.Address) (common.Address, *types.Transaction, *Executor, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorBin), backend, name_, symbol_, nft_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// Executor is an auto generated Go binding around an Ethereum contract.
type Executor struct {
	ExecutorCaller     // Read-only binding to the contract
	ExecutorTransactor // Write-only binding to the contract
	ExecutorFilterer   // Log filterer for contract events
}

// ExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorSession struct {
	Contract     *Executor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorCallerSession struct {
	Contract *ExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorTransactorSession struct {
	Contract     *ExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorRaw struct {
	Contract *Executor // Generic contract binding to access the raw methods on
}

// ExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorCallerRaw struct {
	Contract *ExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorTransactorRaw struct {
	Contract *ExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutor creates a new instance of Executor, bound to a specific deployed contract.
func NewExecutor(address common.Address, backend bind.ContractBackend) (*Executor, error) {
	contract, err := bindExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// NewExecutorCaller creates a new read-only instance of Executor, bound to a specific deployed contract.
func NewExecutorCaller(address common.Address, caller bind.ContractCaller) (*ExecutorCaller, error) {
	contract, err := bindExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorCaller{contract: contract}, nil
}

// NewExecutorTransactor creates a new write-only instance of Executor, bound to a specific deployed contract.
func NewExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorTransactor, error) {
	contract, err := bindExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorTransactor{contract: contract}, nil
}

// NewExecutorFilterer creates a new log filterer instance of Executor, bound to a specific deployed contract.
func NewExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorFilterer, error) {
	contract, err := bindExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorFilterer{contract: contract}, nil
}

// bindExecutor binds a generic wrapper to an already deployed contract.
func bindExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.ExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Executor *ExecutorCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Executor *ExecutorSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Executor.Contract.Allowance(&_Executor.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Executor *ExecutorCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Executor.Contract.Allowance(&_Executor.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Executor *ExecutorCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Executor *ExecutorSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Executor.Contract.BalanceOf(&_Executor.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Executor *ExecutorCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Executor.Contract.BalanceOf(&_Executor.CallOpts, account)
}

// ContributionAmt is a free data retrieval call binding the contract method 0x76c6c2ef.
//
// Solidity: function contributionAmt(address addr) view returns(uint256)
func (_Executor *ExecutorCaller) ContributionAmt(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "contributionAmt", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContributionAmt is a free data retrieval call binding the contract method 0x76c6c2ef.
//
// Solidity: function contributionAmt(address addr) view returns(uint256)
func (_Executor *ExecutorSession) ContributionAmt(addr common.Address) (*big.Int, error) {
	return _Executor.Contract.ContributionAmt(&_Executor.CallOpts, addr)
}

// ContributionAmt is a free data retrieval call binding the contract method 0x76c6c2ef.
//
// Solidity: function contributionAmt(address addr) view returns(uint256)
func (_Executor *ExecutorCallerSession) ContributionAmt(addr common.Address) (*big.Int, error) {
	return _Executor.Contract.ContributionAmt(&_Executor.CallOpts, addr)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Executor *ExecutorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Executor *ExecutorSession) Decimals() (uint8, error) {
	return _Executor.Contract.Decimals(&_Executor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Executor *ExecutorCallerSession) Decimals() (uint8, error) {
	return _Executor.Contract.Decimals(&_Executor.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Executor *ExecutorCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Executor *ExecutorSession) GetBalance() (*big.Int, error) {
	return _Executor.Contract.GetBalance(&_Executor.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Executor *ExecutorCallerSession) GetBalance() (*big.Int, error) {
	return _Executor.Contract.GetBalance(&_Executor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Executor *ExecutorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Executor *ExecutorSession) Name() (string, error) {
	return _Executor.Contract.Name(&_Executor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Executor *ExecutorCallerSession) Name() (string, error) {
	return _Executor.Contract.Name(&_Executor.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Executor *ExecutorCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "onERC721Received", operator, from, tokenId, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Executor *ExecutorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _Executor.Contract.OnERC721Received(&_Executor.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Executor *ExecutorCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _Executor.Contract.OnERC721Received(&_Executor.CallOpts, operator, from, tokenId, data)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Executor *ExecutorCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Executor *ExecutorSession) Price() (*big.Int, error) {
	return _Executor.Contract.Price(&_Executor.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Executor *ExecutorCallerSession) Price() (*big.Int, error) {
	return _Executor.Contract.Price(&_Executor.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Executor *ExecutorCaller) SaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "saleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Executor *ExecutorSession) SaleActive() (bool, error) {
	return _Executor.Contract.SaleActive(&_Executor.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Executor *ExecutorCallerSession) SaleActive() (bool, error) {
	return _Executor.Contract.SaleActive(&_Executor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Executor *ExecutorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Executor *ExecutorSession) Symbol() (string, error) {
	return _Executor.Contract.Symbol(&_Executor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Executor *ExecutorCallerSession) Symbol() (string, error) {
	return _Executor.Contract.Symbol(&_Executor.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Executor *ExecutorCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Executor *ExecutorSession) TotalSupply() (*big.Int, error) {
	return _Executor.Contract.TotalSupply(&_Executor.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Executor *ExecutorCallerSession) TotalSupply() (*big.Int, error) {
	return _Executor.Contract.TotalSupply(&_Executor.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Executor *ExecutorTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Executor *ExecutorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.Approve(&_Executor.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Executor *ExecutorTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.Approve(&_Executor.TransactOpts, spender, amount)
}

// BuyImmediate is a paid mutator transaction binding the contract method 0x5b286a12.
//
// Solidity: function buyImmediate(address receipient) payable returns()
func (_Executor *ExecutorTransactor) BuyImmediate(opts *bind.TransactOpts, receipient common.Address) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "buyImmediate", receipient)
}

// BuyImmediate is a paid mutator transaction binding the contract method 0x5b286a12.
//
// Solidity: function buyImmediate(address receipient) payable returns()
func (_Executor *ExecutorSession) BuyImmediate(receipient common.Address) (*types.Transaction, error) {
	return _Executor.Contract.BuyImmediate(&_Executor.TransactOpts, receipient)
}

// BuyImmediate is a paid mutator transaction binding the contract method 0x5b286a12.
//
// Solidity: function buyImmediate(address receipient) payable returns()
func (_Executor *ExecutorTransactorSession) BuyImmediate(receipient common.Address) (*types.Transaction, error) {
	return _Executor.Contract.BuyImmediate(&_Executor.TransactOpts, receipient)
}

// BuyImmediate0 is a paid mutator transaction binding the contract method 0x6f24539f.
//
// Solidity: function buyImmediate() payable returns()
func (_Executor *ExecutorTransactor) BuyImmediate0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "buyImmediate0")
}

// BuyImmediate0 is a paid mutator transaction binding the contract method 0x6f24539f.
//
// Solidity: function buyImmediate() payable returns()
func (_Executor *ExecutorSession) BuyImmediate0() (*types.Transaction, error) {
	return _Executor.Contract.BuyImmediate0(&_Executor.TransactOpts)
}

// BuyImmediate0 is a paid mutator transaction binding the contract method 0x6f24539f.
//
// Solidity: function buyImmediate() payable returns()
func (_Executor *ExecutorTransactorSession) BuyImmediate0() (*types.Transaction, error) {
	return _Executor.Contract.BuyImmediate0(&_Executor.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0x73e888fd.
//
// Solidity: function contribute(address receipient) payable returns()
func (_Executor *ExecutorTransactor) Contribute(opts *bind.TransactOpts, receipient common.Address) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "contribute", receipient)
}

// Contribute is a paid mutator transaction binding the contract method 0x73e888fd.
//
// Solidity: function contribute(address receipient) payable returns()
func (_Executor *ExecutorSession) Contribute(receipient common.Address) (*types.Transaction, error) {
	return _Executor.Contract.Contribute(&_Executor.TransactOpts, receipient)
}

// Contribute is a paid mutator transaction binding the contract method 0x73e888fd.
//
// Solidity: function contribute(address receipient) payable returns()
func (_Executor *ExecutorTransactorSession) Contribute(receipient common.Address) (*types.Transaction, error) {
	return _Executor.Contract.Contribute(&_Executor.TransactOpts, receipient)
}

// Contribute0 is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Executor *ExecutorTransactor) Contribute0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "contribute0")
}

// Contribute0 is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Executor *ExecutorSession) Contribute0() (*types.Transaction, error) {
	return _Executor.Contract.Contribute0(&_Executor.TransactOpts)
}

// Contribute0 is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Executor *ExecutorTransactorSession) Contribute0() (*types.Transaction, error) {
	return _Executor.Contract.Contribute0(&_Executor.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Executor *ExecutorTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Executor *ExecutorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.DecreaseAllowance(&_Executor.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Executor *ExecutorTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.DecreaseAllowance(&_Executor.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Executor *ExecutorTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Executor *ExecutorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.IncreaseAllowance(&_Executor.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Executor *ExecutorTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.IncreaseAllowance(&_Executor.TransactOpts, spender, addedValue)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 amt) returns()
func (_Executor *ExecutorTransactor) Refund(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "refund", amt)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 amt) returns()
func (_Executor *ExecutorSession) Refund(amt *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.Refund(&_Executor.TransactOpts, amt)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 amt) returns()
func (_Executor *ExecutorTransactorSession) Refund(amt *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.Refund(&_Executor.TransactOpts, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Executor *ExecutorTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Executor *ExecutorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.Transfer(&_Executor.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Executor *ExecutorTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.Transfer(&_Executor.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Executor *ExecutorTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Executor *ExecutorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.TransferFrom(&_Executor.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Executor *ExecutorTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.TransferFrom(&_Executor.TransactOpts, from, to, amount)
}

// ExecutorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Executor contract.
type ExecutorApprovalIterator struct {
	Event *ExecutorApproval // Event containing the contract specifics and raw log

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
func (it *ExecutorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorApproval)
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
		it.Event = new(ExecutorApproval)
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
func (it *ExecutorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorApproval represents a Approval event raised by the Executor contract.
type ExecutorApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Executor *ExecutorFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ExecutorApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorApprovalIterator{contract: _Executor.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Executor *ExecutorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExecutorApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorApproval)
				if err := _Executor.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Executor *ExecutorFilterer) ParseApproval(log types.Log) (*ExecutorApproval, error) {
	event := new(ExecutorApproval)
	if err := _Executor.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Executor contract.
type ExecutorTransferIterator struct {
	Event *ExecutorTransfer // Event containing the contract specifics and raw log

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
func (it *ExecutorTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorTransfer)
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
		it.Event = new(ExecutorTransfer)
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
func (it *ExecutorTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorTransfer represents a Transfer event raised by the Executor contract.
type ExecutorTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Executor *ExecutorFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorTransferIterator{contract: _Executor.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Executor *ExecutorFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExecutorTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorTransfer)
				if err := _Executor.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Executor *ExecutorFilterer) ParseTransfer(log types.Log) (*ExecutorTransfer, error) {
	event := new(ExecutorTransfer)
	if err := _Executor.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
