package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// An account stores details of an ethereum account
type Account struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
	WalletAddr common.Address
	Nonce      uint64
}

// An executor launches smart contracts
type Executor struct {
	Address string
	Client  *ethclient.Client
	Txs     *bind.TransactOpts
	Account
}

// A contract is an interface for a solidity-generated
// contract package
type ERC721Contract func(txs *bind.TransactOpts, client *ethclient.Client,
	name string, symbol string, price *big.Int, addr common.Address) (common.Address,
	*types.Transaction, error)

type ERC20ExecutorContract func(txs *bind.TransactOpts, client *ethclient.Client,
	name string, symbol string, addr common.Address) (common.Address,
	*types.Transaction, error)

// Create an executor. 'address' is the server address of the
// ethereum node.
func NewExecutor(privKey string, address string) (exec Executor) {
	client, err := ethclient.Dial(address)
	if err != nil {
		panic(fmt.Sprintf("Failed to dial eth client at address %s", address))
	}

	exec.Client = client
	exec.Address = address

	// Create transactor
	ethKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		fmt.Errorf("Invalid private key")
	}

	pubKey := ethKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Errorf("Invalid private key")
	}

	walletAddr := crypto.PubkeyToAddress(*pubKeyECDSA)
	nonce, err := exec.Client.PendingNonceAt(context.Background(), walletAddr)
	if err != nil {
		fmt.Errorf("Unable to retrieve latest address nonce: %s", err)
	}

	chainID, err := exec.Client.ChainID(context.Background())
	if err != nil {
		fmt.Errorf("Unable to retrieve Chain ID: %s", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ethKey, chainID)
	if err != nil {
		panic(err)
	}

	// Finalize smart contract executor
	exec.PrivateKey = *ethKey
	exec.PublicKey = *pubKeyECDSA
	exec.WalletAddr = walletAddr
	exec.Nonce = nonce
	exec.Txs = auth

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(875000000)

	return exec
}

// Launch the ERC721 sale contract to the blockchain
func (exec Executor) DeployERC721Contract(cn ERC721Contract, name string,
	price *big.Int, addr common.Address) (common.Address, *types.Transaction,
	error) {

	address, tx, err := cn(exec.Txs, exec.Client, name, "CNFT", price, addr)
	return address, tx, err
}

// Launch the ERC20 Executor contract to the blockchain
func (exec Executor) DeployERC20ExecutorContract(cn ERC20ExecutorContract,
	name string, addr common.Address) (common.Address, *types.Transaction, error) {

	address, tx, err := cn(exec.Txs, exec.Client, name, "CNFT", addr)
	return address, tx, err
}
