package routes

import (
	"net/http"
	"strconv"

	"github.com/Cameronketchem/CEN3031-Group91/server/data"
	"github.com/gin-gonic/gin"
)

type routeHandler func(store *data.Store, c *gin.Context)

// api/user/:id
func getUserById(store *data.Store, c *gin.Context) {
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
func getAssetById(store *data.Store, c *gin.Context) {
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
func getContractById(store *data.Store, c *gin.Context) {
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
func getBatchUsers(store *data.Store, c *gin.Context) {
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
func getBatchAssets(store *data.Store, c *gin.Context) {
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

	for i := 0; i < 20; i++ {
		EContract, Eerr := store.QueryContractById(assets[i].ExecContract)
		AContract, Aerr := store.QueryContractById(assets[i].AssetContract)
		Executor, _ := store.QueryUserById(assets[i].Executor)

		if Eerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No exec contract found for asset ID " + strconv.Itoa(lastID+i)})
			return
		}

		if Aerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No asset contract found for asset ID " + strconv.Itoa(lastID+i)})
			return
		}

		JSONassets = append(JSONassets, gin.H{"asset_id": assets[i].AssetID, "exec_contract": EContract, "asset_contract": AContract, "img_preview": assets[i].ImgPreview, "executor": Executor, "desc": assets[i].Desc, "price": assets[i].Price})
	}
	c.JSON(http.StatusOK, JSONassets)
}
