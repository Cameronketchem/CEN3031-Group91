// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721sale

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

// Erc721saleMetaData contains all meta data concerning the Erc721sale contract.
var Erc721saleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sale_receipient_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001c4138038062001c4183398101604081905262000034916200016a565b838360006200004483826200028e565b5060016200005382826200028e565b50505060089190915560068054600160a01b6001600160a81b03199091163360ff60a01b191617179055600780546001600160a01b0319166001600160a01b03909216919091179055506200035a9050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000cd57600080fd5b81516001600160401b0380821115620000ea57620000ea620000a5565b604051601f8301601f19908116603f01168101908282118183101715620001155762000115620000a5565b816040528381526020925086838588010111156200013257600080fd5b600091505b8382101562000156578582018301518183018401529082019062000137565b600093810190920192909252949350505050565b600080600080608085870312156200018157600080fd5b84516001600160401b03808211156200019957600080fd5b620001a788838901620000bb565b95506020870151915080821115620001be57600080fd5b50620001cd87828801620000bb565b60408701516060880151919550935090506001600160a01b0381168114620001f457600080fd5b939692955090935050565b600181811c908216806200021457607f821691505b6020821081036200023557634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028957600081815260208120601f850160051c81016020861015620002645750805b601f850160051c820191505b81811015620002855782815560010162000270565b5050505b505050565b81516001600160401b03811115620002aa57620002aa620000a5565b620002c281620002bb8454620001ff565b846200023b565b602080601f831160018114620002fa5760008415620002e15750858301515b600019600386901b1c1916600185901b17855562000285565b600085815260208120601f198616915b828110156200032b578886015182559484019460019091019084016200030a565b50858210156200034a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6118d7806200036a6000396000f3fe60806040526004361061011f5760003560e01c806368428a1b116100a0578063a2b40d1911610064578063a2b40d191461030b578063b88d4fde1461032b578063c87b56dd1461034b578063e985e9c51461036b578063f088d5471461038b57600080fd5b806368428a1b1461027e57806370a082311461029357806395d89b41146102c1578063a035b1fe146102d6578063a22cb465146102eb57600080fd5b806333e364cb116100e757806333e364cb146101f557806342842e0e1461020a57806355367ba91461022a5780635c975abb1461023f5780636352211e1461025e57600080fd5b806301ffc9a71461012457806306fdde0314610159578063081812fc1461017b578063095ea7b3146101b357806323b872dd146101d5575b600080fd5b34801561013057600080fd5b5061014461013f366004611427565b61039e565b60405190151581526020015b60405180910390f35b34801561016557600080fd5b5061016e6103f0565b6040516101509190611494565b34801561018757600080fd5b5061019b6101963660046114a7565b610482565b6040516001600160a01b039091168152602001610150565b3480156101bf57600080fd5b506101d36101ce3660046114dc565b6104a9565b005b3480156101e157600080fd5b506101d36101f0366004611506565b6105c3565b34801561020157600080fd5b506101d36105f4565b34801561021657600080fd5b506101d3610225366004611506565b6106d9565b34801561023657600080fd5b506101d36106f4565b34801561024b57600080fd5b50600654600160a81b900460ff16610144565b34801561026a57600080fd5b5061019b6102793660046114a7565b6107cc565b34801561028a57600080fd5b5061014461082c565b34801561029f57600080fd5b506102b36102ae366004611542565b610857565b604051908152602001610150565b3480156102cd57600080fd5b5061016e6108dd565b3480156102e257600080fd5b506008546102b3565b3480156102f757600080fd5b506101d361030636600461155d565b6108ec565b34801561031757600080fd5b506101d36103263660046114a7565b6108fb565b34801561033757600080fd5b506101d36103463660046115af565b6109b8565b34801561035757600080fd5b5061016e6103663660046114a7565b6109f0565b34801561037757600080fd5b5061014461038636600461168b565b610a64565b610144610399366004611542565b610a92565b60006001600160e01b031982166380ac58cd60e01b14806103cf57506001600160e01b03198216635b5e139f60e01b145b806103ea57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600080546103ff906116be565b80601f016020809104026020016040519081016040528092919081815260200182805461042b906116be565b80156104785780601f1061044d57610100808354040283529160200191610478565b820191906000526020600020905b81548152906001019060200180831161045b57829003601f168201915b5050505050905090565b600061048d82610b96565b506000908152600460205260409020546001600160a01b031690565b60006104b4826107cc565b9050806001600160a01b0316836001600160a01b0316036105265760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061054257506105428133610a64565b6105b45760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000606482015260840161051d565b6105be8383610bf8565b505050565b6105cd3382610c66565b6105e95760405162461bcd60e51b815260040161051d906116f8565b6105be838383610cc5565b600654600160a81b900460ff1661060757565b61060f61082c565b61065b5760405162461bcd60e51b815260206004820152601f60248201527f43616e6e6f7420756e706175736520616e20696e6163746976652073616c6500604482015260640161051d565b6006546001600160a01b031633148061067e57506007546001600160a01b031633145b6106ca5760405162461bcd60e51b815260206004820152601b60248201527f4f6e6c792061646d696e2063616e20756e70617573652073616c650000000000604482015260640161051d565b6006805460ff60a81b19169055565b6105be838383604051806020016040528060008152506109b8565b6106fc61082c565b6107485760405162461bcd60e51b815260206004820152601d60248201527f43616e6e6f7420706175736520616e20696e6163746976652073616c65000000604482015260640161051d565b6006546001600160a01b031633148061076b57506007546001600160a01b031633145b6107b75760405162461bcd60e51b815260206004820152601960248201527f4f6e6c792061646d696e2063616e2070617573652073616c6500000000000000604482015260640161051d565b6006805460ff60a81b1916600160a81b179055565b6000818152600260205260408120546001600160a01b0316806103ea5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161051d565b600654600090600160a01b900460ff1680156108525750600654600160a81b900460ff16155b905090565b60006001600160a01b0382166108c15760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161051d565b506001600160a01b031660009081526003602052604090205490565b6060600180546103ff906116be565b6108f7338383610e36565b5050565b61090361082c565b6109445760405162461bcd60e51b815260206004820152601260248201527153616c65206973206e6f742061637469766560701b604482015260640161051d565b6006546001600160a01b031633148061096757506007546001600160a01b031633145b6109b35760405162461bcd60e51b815260206004820152601b60248201527f4f6e6c792061646d696e2063616e206368616e67652070726963650000000000604482015260640161051d565b600855565b6109c23383610c66565b6109de5760405162461bcd60e51b815260040161051d906116f8565b6109ea84848484610f04565b50505050565b60606109fb82610b96565b6000610a1260408051602081019091526000815290565b90506000815111610a325760405180602001604052806000815250610a5d565b80610a3c84610f37565b604051602001610a4d929190611745565b6040516020818303038152906040525b9392505050565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6000610a9c61082c565b610add5760405162461bcd60e51b815260206004820152601260248201527153616c65206973206e6f742061637469766560701b604482015260640161051d565b600854341015610b3d5760405162461bcd60e51b815260206004820152602560248201527f53656e742076616c7565206d75737420657175616c206f722065786365656420604482015264707269636560d81b606482015260840161051d565b6007546040516001600160a01b03909116903480156108fc02916000818181858888f19350505050158015610b76573d6000803e3d6000fd5b50610b82826000610fca565b50506006805460ff60a01b19169055600190565b6000818152600260205260409020546001600160a01b0316610bf55760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161051d565b50565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610c2d826107cc565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610c72836107cc565b9050806001600160a01b0316846001600160a01b03161480610c995750610c998185610a64565b80610cbd5750836001600160a01b0316610cb284610482565b6001600160a01b0316145b949350505050565b826001600160a01b0316610cd8826107cc565b6001600160a01b031614610cfe5760405162461bcd60e51b815260040161051d90611774565b6001600160a01b038216610d605760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161051d565b610d6d8383836001610fe4565b826001600160a01b0316610d80826107cc565b6001600160a01b031614610da65760405162461bcd60e51b815260040161051d90611774565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b816001600160a01b0316836001600160a01b031603610e975760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161051d565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610f0f848484610cc5565b610f1b8484848461106c565b6109ea5760405162461bcd60e51b815260040161051d906117b9565b60606000610f448361116d565b600101905060008167ffffffffffffffff811115610f6457610f64611599565b6040519080825280601f01601f191660200182016040528015610f8e576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610f9857509392505050565b6108f7828260405180602001604052806000815250611245565b60018111156109ea576001600160a01b0384161561102a576001600160a01b03841660009081526003602052604081208054839290611024908490611821565b90915550505b6001600160a01b038316156109ea576001600160a01b03831660009081526003602052604081208054839290611061908490611834565b909155505050505050565b60006001600160a01b0384163b1561116257604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906110b0903390899088908890600401611847565b6020604051808303816000875af19250505080156110eb575060408051601f3d908101601f191682019092526110e891810190611884565b60015b611148573d808015611119576040519150601f19603f3d011682016040523d82523d6000602084013e61111e565b606091505b5080516000036111405760405162461bcd60e51b815260040161051d906117b9565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610cbd565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106111ac5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106111d8576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106111f657662386f26fc10000830492506010015b6305f5e100831061120e576305f5e100830492506008015b612710831061122257612710830492506004015b60648310611234576064830492506002015b600a83106103ea5760010192915050565b61124f8383611278565b61125c600084848461106c565b6105be5760405162461bcd60e51b815260040161051d906117b9565b6001600160a01b0382166112ce5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161051d565b6000818152600260205260409020546001600160a01b0316156113335760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161051d565b611341600083836001610fe4565b6000818152600260205260409020546001600160a01b0316156113a65760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161051d565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6001600160e01b031981168114610bf557600080fd5b60006020828403121561143957600080fd5b8135610a5d81611411565b60005b8381101561145f578181015183820152602001611447565b50506000910152565b60008151808452611480816020860160208601611444565b601f01601f19169290920160200192915050565b602081526000610a5d6020830184611468565b6000602082840312156114b957600080fd5b5035919050565b80356001600160a01b03811681146114d757600080fd5b919050565b600080604083850312156114ef57600080fd5b6114f8836114c0565b946020939093013593505050565b60008060006060848603121561151b57600080fd5b611524846114c0565b9250611532602085016114c0565b9150604084013590509250925092565b60006020828403121561155457600080fd5b610a5d826114c0565b6000806040838503121561157057600080fd5b611579836114c0565b91506020830135801515811461158e57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600080600080608085870312156115c557600080fd5b6115ce856114c0565b93506115dc602086016114c0565b925060408501359150606085013567ffffffffffffffff8082111561160057600080fd5b818701915087601f83011261161457600080fd5b81358181111561162657611626611599565b604051601f8201601f19908116603f0116810190838211818310171561164e5761164e611599565b816040528281528a602084870101111561166757600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561169e57600080fd5b6116a7836114c0565b91506116b5602084016114c0565b90509250929050565b600181811c908216806116d257607f821691505b6020821081036116f257634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b60008351611757818460208801611444565b83519083019061176b818360208801611444565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b818103818111156103ea576103ea61180b565b808201808211156103ea576103ea61180b565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061187a90830184611468565b9695505050505050565b60006020828403121561189657600080fd5b8151610a5d8161141156fea264697066735822122062f1f34bb042453e31ba7746ff9d7c64ae9057feef41f074413ba6bfcf804bff64736f6c63430008120033",
}

// Erc721saleABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721saleMetaData.ABI instead.
var Erc721saleABI = Erc721saleMetaData.ABI

// Erc721saleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc721saleMetaData.Bin instead.
var Erc721saleBin = Erc721saleMetaData.Bin

// DeployErc721sale deploys a new Ethereum contract, binding an instance of Erc721sale to it.
func DeployErc721sale(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, price_ *big.Int, sale_receipient_ common.Address) (common.Address, *types.Transaction, *Erc721sale, error) {
	parsed, err := Erc721saleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc721saleBin), backend, name_, symbol_, price_, sale_receipient_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc721sale{Erc721saleCaller: Erc721saleCaller{contract: contract}, Erc721saleTransactor: Erc721saleTransactor{contract: contract}, Erc721saleFilterer: Erc721saleFilterer{contract: contract}}, nil
}

// Erc721sale is an auto generated Go binding around an Ethereum contract.
type Erc721sale struct {
	Erc721saleCaller     // Read-only binding to the contract
	Erc721saleTransactor // Write-only binding to the contract
	Erc721saleFilterer   // Log filterer for contract events
}

// Erc721saleCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721saleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721saleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721saleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721saleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721saleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721saleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721saleSession struct {
	Contract     *Erc721sale       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721saleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721saleCallerSession struct {
	Contract *Erc721saleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Erc721saleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721saleTransactorSession struct {
	Contract     *Erc721saleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Erc721saleRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721saleRaw struct {
	Contract *Erc721sale // Generic contract binding to access the raw methods on
}

// Erc721saleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721saleCallerRaw struct {
	Contract *Erc721saleCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721saleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721saleTransactorRaw struct {
	Contract *Erc721saleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721sale creates a new instance of Erc721sale, bound to a specific deployed contract.
func NewErc721sale(address common.Address, backend bind.ContractBackend) (*Erc721sale, error) {
	contract, err := bindErc721sale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721sale{Erc721saleCaller: Erc721saleCaller{contract: contract}, Erc721saleTransactor: Erc721saleTransactor{contract: contract}, Erc721saleFilterer: Erc721saleFilterer{contract: contract}}, nil
}

// NewErc721saleCaller creates a new read-only instance of Erc721sale, bound to a specific deployed contract.
func NewErc721saleCaller(address common.Address, caller bind.ContractCaller) (*Erc721saleCaller, error) {
	contract, err := bindErc721sale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721saleCaller{contract: contract}, nil
}

// NewErc721saleTransactor creates a new write-only instance of Erc721sale, bound to a specific deployed contract.
func NewErc721saleTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721saleTransactor, error) {
	contract, err := bindErc721sale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721saleTransactor{contract: contract}, nil
}

// NewErc721saleFilterer creates a new log filterer instance of Erc721sale, bound to a specific deployed contract.
func NewErc721saleFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721saleFilterer, error) {
	contract, err := bindErc721sale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721saleFilterer{contract: contract}, nil
}

// bindErc721sale binds a generic wrapper to an already deployed contract.
func bindErc721sale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721saleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721sale *Erc721saleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721sale.Contract.Erc721saleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721sale *Erc721saleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721sale.Contract.Erc721saleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721sale *Erc721saleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721sale.Contract.Erc721saleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721sale *Erc721saleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721sale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721sale *Erc721saleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721sale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721sale *Erc721saleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721sale.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721sale *Erc721saleCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721sale *Erc721saleSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721sale.Contract.BalanceOf(&_Erc721sale.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721sale *Erc721saleCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721sale.Contract.BalanceOf(&_Erc721sale.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721sale *Erc721saleCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721sale *Erc721saleSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721sale.Contract.GetApproved(&_Erc721sale.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721sale *Erc721saleCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721sale.Contract.GetApproved(&_Erc721sale.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721sale *Erc721saleCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721sale *Erc721saleSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721sale.Contract.IsApprovedForAll(&_Erc721sale.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721sale *Erc721saleCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721sale.Contract.IsApprovedForAll(&_Erc721sale.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721sale *Erc721saleCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721sale *Erc721saleSession) Name() (string, error) {
	return _Erc721sale.Contract.Name(&_Erc721sale.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721sale *Erc721saleCallerSession) Name() (string, error) {
	return _Erc721sale.Contract.Name(&_Erc721sale.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721sale *Erc721saleCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721sale *Erc721saleSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721sale.Contract.OwnerOf(&_Erc721sale.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721sale *Erc721saleCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721sale.Contract.OwnerOf(&_Erc721sale.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc721sale *Erc721saleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc721sale *Erc721saleSession) Paused() (bool, error) {
	return _Erc721sale.Contract.Paused(&_Erc721sale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc721sale *Erc721saleCallerSession) Paused() (bool, error) {
	return _Erc721sale.Contract.Paused(&_Erc721sale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Erc721sale *Erc721saleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Erc721sale *Erc721saleSession) Price() (*big.Int, error) {
	return _Erc721sale.Contract.Price(&_Erc721sale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Erc721sale *Erc721saleCallerSession) Price() (*big.Int, error) {
	return _Erc721sale.Contract.Price(&_Erc721sale.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Erc721sale *Erc721saleCaller) SaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "saleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Erc721sale *Erc721saleSession) SaleActive() (bool, error) {
	return _Erc721sale.Contract.SaleActive(&_Erc721sale.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Erc721sale *Erc721saleCallerSession) SaleActive() (bool, error) {
	return _Erc721sale.Contract.SaleActive(&_Erc721sale.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721sale *Erc721saleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721sale *Erc721saleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721sale.Contract.SupportsInterface(&_Erc721sale.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721sale *Erc721saleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721sale.Contract.SupportsInterface(&_Erc721sale.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721sale *Erc721saleCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721sale *Erc721saleSession) Symbol() (string, error) {
	return _Erc721sale.Contract.Symbol(&_Erc721sale.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721sale *Erc721saleCallerSession) Symbol() (string, error) {
	return _Erc721sale.Contract.Symbol(&_Erc721sale.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721sale *Erc721saleCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Erc721sale.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721sale *Erc721saleSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721sale.Contract.TokenURI(&_Erc721sale.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721sale *Erc721saleCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721sale.Contract.TokenURI(&_Erc721sale.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.Contract.Approve(&_Erc721sale.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.Contract.Approve(&_Erc721sale.TransactOpts, to, tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address buyer) payable returns(bool)
func (_Erc721sale *Erc721saleTransactor) Buy(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "buy", buyer)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address buyer) payable returns(bool)
func (_Erc721sale *Erc721saleSession) Buy(buyer common.Address) (*types.Transaction, error) {
	return _Erc721sale.Contract.Buy(&_Erc721sale.TransactOpts, buyer)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address buyer) payable returns(bool)
func (_Erc721sale *Erc721saleTransactorSession) Buy(buyer common.Address) (*types.Transaction, error) {
	return _Erc721sale.Contract.Buy(&_Erc721sale.TransactOpts, buyer)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 price_) returns()
func (_Erc721sale *Erc721saleTransactor) ChangePrice(opts *bind.TransactOpts, price_ *big.Int) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "changePrice", price_)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 price_) returns()
func (_Erc721sale *Erc721saleSession) ChangePrice(price_ *big.Int) (*types.Transaction, error) {
	return _Erc721sale.Contract.ChangePrice(&_Erc721sale.TransactOpts, price_)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 price_) returns()
func (_Erc721sale *Erc721saleTransactorSession) ChangePrice(price_ *big.Int) (*types.Transaction, error) {
	return _Erc721sale.Contract.ChangePrice(&_Erc721sale.TransactOpts, price_)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_Erc721sale *Erc721saleTransactor) PauseSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "pauseSale")
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_Erc721sale *Erc721saleSession) PauseSale() (*types.Transaction, error) {
	return _Erc721sale.Contract.PauseSale(&_Erc721sale.TransactOpts)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_Erc721sale *Erc721saleTransactorSession) PauseSale() (*types.Transaction, error) {
	return _Erc721sale.Contract.PauseSale(&_Erc721sale.TransactOpts)
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_Erc721sale *Erc721saleTransactor) ResumeSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "resumeSale")
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_Erc721sale *Erc721saleSession) ResumeSale() (*types.Transaction, error) {
	return _Erc721sale.Contract.ResumeSale(&_Erc721sale.TransactOpts)
}

// ResumeSale is a paid mutator transaction binding the contract method 0x33e364cb.
//
// Solidity: function resumeSale() returns()
func (_Erc721sale *Erc721saleTransactorSession) ResumeSale() (*types.Transaction, error) {
	return _Erc721sale.Contract.ResumeSale(&_Erc721sale.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.Contract.SafeTransferFrom(&_Erc721sale.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.Contract.SafeTransferFrom(&_Erc721sale.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc721sale *Erc721saleTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc721sale *Erc721saleSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc721sale.Contract.SafeTransferFrom0(&_Erc721sale.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc721sale *Erc721saleTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc721sale.Contract.SafeTransferFrom0(&_Erc721sale.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721sale *Erc721saleTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721sale *Erc721saleSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721sale.Contract.SetApprovalForAll(&_Erc721sale.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721sale *Erc721saleTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721sale.Contract.SetApprovalForAll(&_Erc721sale.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.Contract.TransferFrom(&_Erc721sale.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721sale *Erc721saleTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721sale.Contract.TransferFrom(&_Erc721sale.TransactOpts, from, to, tokenId)
}

// Erc721saleApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc721sale contract.
type Erc721saleApprovalIterator struct {
	Event *Erc721saleApproval // Event containing the contract specifics and raw log

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
func (it *Erc721saleApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721saleApproval)
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
		it.Event = new(Erc721saleApproval)
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
func (it *Erc721saleApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721saleApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721saleApproval represents a Approval event raised by the Erc721sale contract.
type Erc721saleApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721sale *Erc721saleFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Erc721saleApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721sale.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721saleApprovalIterator{contract: _Erc721sale.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721sale *Erc721saleFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc721saleApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721sale.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721saleApproval)
				if err := _Erc721sale.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721sale *Erc721saleFilterer) ParseApproval(log types.Log) (*Erc721saleApproval, error) {
	event := new(Erc721saleApproval)
	if err := _Erc721sale.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721saleApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc721sale contract.
type Erc721saleApprovalForAllIterator struct {
	Event *Erc721saleApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Erc721saleApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721saleApprovalForAll)
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
		it.Event = new(Erc721saleApprovalForAll)
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
func (it *Erc721saleApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721saleApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721saleApprovalForAll represents a ApprovalForAll event raised by the Erc721sale contract.
type Erc721saleApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721sale *Erc721saleFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Erc721saleApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721sale.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc721saleApprovalForAllIterator{contract: _Erc721sale.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721sale *Erc721saleFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc721saleApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721sale.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721saleApprovalForAll)
				if err := _Erc721sale.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721sale *Erc721saleFilterer) ParseApprovalForAll(log types.Log) (*Erc721saleApprovalForAll, error) {
	event := new(Erc721saleApprovalForAll)
	if err := _Erc721sale.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721saleTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc721sale contract.
type Erc721saleTransferIterator struct {
	Event *Erc721saleTransfer // Event containing the contract specifics and raw log

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
func (it *Erc721saleTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721saleTransfer)
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
		it.Event = new(Erc721saleTransfer)
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
func (it *Erc721saleTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721saleTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721saleTransfer represents a Transfer event raised by the Erc721sale contract.
type Erc721saleTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721sale *Erc721saleFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Erc721saleTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721sale.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721saleTransferIterator{contract: _Erc721sale.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721sale *Erc721saleFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc721saleTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721sale.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721saleTransfer)
				if err := _Erc721sale.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721sale *Erc721saleFilterer) ParseTransfer(log types.Log) (*Erc721saleTransfer, error) {
	event := new(Erc721saleTransfer)
	if err := _Erc721sale.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
