package executor

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Implements Contract type in ../contracts.go
func ERC20EXECUTOR(txs *bind.TransactOpts, client *ethclient.Client,
	name string, symbol string, addr common.Address) (common.Address,
	*types.Transaction, error) {

	address, tx, _, err := DeployExecutor(txs, client, name, symbol, addr)
	return address, tx, err
}
