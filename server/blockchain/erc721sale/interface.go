package erc721sale

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// Implements Contract type in ../contracts.go
func ERC721SALE(txs *bind.TransactOpts, client *ethclient.Client,
	name string, symbol string, price *big.Int, addr common.Address) (common.Address,
	*types.Transaction, error) {

	address, tx, _, err := DeployErc721sale(txs, client, name, symbol, price, addr)
	return address, tx, err
}
