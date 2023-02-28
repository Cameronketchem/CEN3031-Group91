BUILD_DIR = ../../build

ERC20Executor:
	solc --optimize --overwrite --abi ./solidity/ERC20Executor.sol -o $(BUILD_DIR)
	solc --optimize --overwrite --bin ./solidity/ERC20Executor.sol -o $(BUILD_DIR)
	abigen --abi=$(BUILD_DIR)/ERC20Executor.abi --bin=$(BUILD_DIR)/ERC20Executor.bin --pkg=executor --out=./executor/executor.go
	solc --optimize --overwrite --abi ./solidity/ERC721Sale.sol -o $(BUILD_DIR)
	solc --optimize --overwrite --bin ./solidity/ERC721Sale.sol -o $(BUILD_DIR)
	abigen --abi=$(BUILD_DIR)/ERC721Sale.abi --bin=$(BUILD_DIR)/ERC721Sale.bin --pkg=erc721sale --out=./erc721sale/erc721sale.go
