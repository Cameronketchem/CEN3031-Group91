package routes

import (
	"fmt"
	"github.com/Cameronketchem/CEN3031-Group91/server/blockchain"
	"github.com/Cameronketchem/CEN3031-Group91/server/blockchain/erc721sale"
	"github.com/Cameronketchem/CEN3031-Group91/server/blockchain/executor"
	"github.com/Cameronketchem/CEN3031-Group91/server/data"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

// A route handler is a middleware function, with access to API
// struct internals, that gets passed to the API wrapper functions
// in api.go.
type routeHandler func(store *data.Store, cont *blockchain.Executor,
	c *gin.Context)

// api/user/:id
func getUserById(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	user, err := store.QueryUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No users with that ID"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// api/asset/:id
func getAssetById(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	asset, err := store.QueryAssetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No assets with that ID"})
		return
	}

	EContract, Eerr := store.QueryContractById(asset.ExecContract)
	AContract, Aerr := store.QueryContractById(asset.AssetContract)
	Executor, _ := store.QueryUserById(asset.Executor)

	if Eerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No exec contract found."})
		return
	}

	if Aerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No asset contract found"})
		return
	}

	JSONAsset := gin.H{"asset_id": asset.AssetID, "exec_contract": EContract, "asset_contract": AContract, "img_preview": asset.ImgPreview, "executor": Executor, "desc": asset.Desc, "price": asset.Price}
	c.JSON(http.StatusOK, JSONAsset)

}

// api/contract/:id
func getContractById(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	contract, err := store.QueryContractById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No contracts with that ID"})
		return
	}

	c.JSON(http.StatusOK, contract)
}

// api/users/:lastID
func getBatchUsers(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	lastID, err := strconv.Atoi(c.Params.ByName("lastID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	users, err := store.GetUsers(lastID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No users with that ID"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// api/assets/:lastID
func getBatchAssets(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	lastID, err := strconv.Atoi(c.Params.ByName("lastID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	assets, err := store.GetAssets(lastID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No assets with that ID"})
		return
	}

	var JSONassets = []gin.H{}

	for i := 0; i < len(assets); i++ {
		EContract, Eerr := store.QueryContractById(assets[i].ExecContract)
		AContract, Aerr := store.QueryContractById(assets[i].AssetContract)
		Executor, _ := store.QueryUserById(assets[i].Executor)

		if Eerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("No exec contract found for AssetID:%d",
					lastID+i)})
			return
		}

		if Aerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("No asset contract found for AssetID:%d",
					lastID+i)})
			return
		}

		JSONassets = append(JSONassets, gin.H{
			"asset_id":       assets[i].AssetID,
			"exec_contract":  EContract,
			"asset_contract": AContract,
			"img_preview":    assets[i].ImgPreview,
			"executor":       Executor,
			"desc":           assets[i].Desc,
			"price":          assets[i].Price,
		})
	}

	c.JSON(http.StatusOK, JSONassets)
}

// api/asset
type postAssetInput struct {
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float64 `json:"price"`
	ImgPreview string  `json:"img_preview"`
}

// api/asset
func postAsset(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	// Validate input
	var input postAssetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Create asset
	var asset data.Asset
	asset.ImgPreview = input.ImgPreview
	asset.Price = input.Price
	asset.Desc = input.Desc
	asset.Executor = c.Value("jwt-userid").(int)

	// Price is in ether. Needs to be converted into wei integer.
	priceflt := big.NewFloat(input.Price)
	factor := big.NewFloat(1000000000000000000) // 10^18
	priceflt.Mul(priceflt, factor)

	price := new(big.Int)
	priceflt.Int(price)

	// Address needs to be in common.Address
	address := common.HexToAddress(c.Value("jwt-address").(string))

	// Create asset contract.
	// The owner is the authenticated user.
	assetAddr, _, err := cont.DeployERC721Contract(erc721sale.ERC721SALE, input.Name,
		price, address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Asset creation service currently unavailable"})
		return
	}

	assetContract := data.Contract{0, assetAddr.Hex(), false, false}
	if err := store.InsertContract(assetContract); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert asseet contract to db"})
		return
	}

	// Create exeuctor contract.
	// The owner is the API executor's address.
	execAddr, _, err := cont.DeployERC20ExecutorContract(executor.ERC20EXECUTOR,
		input.Name, common.HexToAddress(cont.Address))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Asset creation service currently unavailable"})
		return
	}

	execContract := data.Contract{0, execAddr.Hex(), true, false}
	if err := store.InsertContract(execContract); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert exec contract to db"})
		return
	}

	// Finalize asset.
	assetContractDb, err := store.QueryContractByAddr(assetAddr.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save asset contract in db"})
		return
	}

	execContractDb, err := store.QueryContractByAddr(execAddr.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save exec contract to db"})
		return
	}

	asset.ExecContract = execContractDb.ContractID
	asset.AssetContract = assetContractDb.ContractID

	// Insert to database.
	if err := store.InsertAsset(asset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not insert to database, try again later"})
		return
	}

	assetJson, _ := asset.ToJSON()
	c.JSON(http.StatusCreated, gin.H{"asset": assetJson})
}

// api/asset/:id/price
func getAssetPrice(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	// Get ID as integer from string.
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	// Verify asset exists.
	asset, err := store.QueryAssetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No assets with that ID"})
		return
	}

	// Get asset price.
	assetCnt, err := store.QueryContractById(asset.AssetContract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot find the asset's assosciated contract"})
		return
	}

	addr := assetCnt.Addr
	contex, err := erc721sale.NewErc721sale(common.HexToAddress(addr), cont.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot reach asset contract"})
		return
	}

	price, err := contex.Price(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot reach asset contract"})
		return
	}

	factor := big.NewInt(1000000000000000000) // 10^18
	price.Quo(price, factor)

	c.JSON(http.StatusOK, price.Uint64())
}

// api/asset/:id/active
func getAssetActive(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	// Get ID as integer from string.
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	// Verify asset exists.
	asset, err := store.QueryAssetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No assets with that ID"})
		return
	}

	// Get asset state.
	assetCnt, err := store.QueryContractById(asset.ExecContract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot find the asset's assosciated contract"})
		return
	}

	addr := assetCnt.Addr
	contex, err := executor.NewExecutorCaller(common.HexToAddress(addr), cont.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot reach asset contract"})
		return
	}

	active, err := contex.SaleActive(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot retrieve asset state"})
		return
	}

	c.JSON(http.StatusOK, active)
}

// api/asset/:id/contrib/:addr
func getAssetContribByAddr(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	// Get ID as integer from string.
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	// Verify asset exists.
	asset, err := store.QueryAssetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No assets with that ID"})
		return
	}

	// Get address parameter as common.Address type.
	address := common.HexToAddress(c.Params.ByName("addr"))

	// Get address contributions.
	assetCnt, err := store.QueryContractById(asset.ExecContract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot find the asset's assosciated contract"})
		return
	}

	contractAddr := assetCnt.Addr
	contex, err := executor.NewExecutorCaller(common.HexToAddress(contractAddr), cont.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot reach asset contract"})
		return
	}

	contrib, err := contex.ContributionAmt(&bind.CallOpts{}, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot retrieve asset contributions"})
		return
	}

	// Convert contrib to ETH/int64
	factor := big.NewInt(1000000000000000000) // 10^18
	contrib.Quo(contrib, factor)

	c.JSON(http.StatusOK, contrib.Uint64())
}

// api/asset/:id/balance
func getAssetBalance(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	// Get ID as integer from string.
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter expected integer"})
		return
	}

	// Verify asset exists.
	asset, err := store.QueryAssetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No assets with that ID"})
		return
	}

	// Get asset contract balance.
	assetCnt, err := store.QueryContractById(asset.ExecContract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot find the asset's assosciated contract"})
		return
	}

	addr := assetCnt.Addr
	contex, err := executor.NewExecutorCaller(common.HexToAddress(addr), cont.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot reach asset contract"})
		return
	}

	balance, err := contex.GetBalance(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot retrieve asset balance"})
		return
	}

	// Convert to ETH/uint64.
	factor := big.NewInt(1000000000000000000) // 10^18
	balance.Quo(balance, factor)

	c.JSON(http.StatusOK, balance.Uint64())
}
