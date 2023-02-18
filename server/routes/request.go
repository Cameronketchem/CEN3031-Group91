package routes

import (
	"github.com/Cameronketchem/CEN3031-Group91/server/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	c.JSON(http.StatusOK, asset)
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

	c.JSON(http.StatusOK, assets)
}
