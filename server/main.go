package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Cameronketchem/CEN3031-Group91/server/data"
	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store, err := data.OpenStore("fileName.sqlite3", false)
	if err != nil {
		fmt.Errorf("Failed to open store\n")
	}

	store.Init(db.SynchronousModeFull, true)
	store.Seed()
	defer store.Close()

	//Return user with ID passed into the function
	r.GET("api/user/:id", func(c *gin.Context) {

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
	})

	//Return asset with ID passed into the function
	r.GET("api/asset/:id", func(c *gin.Context) {

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
	})

	//Return contract with ID passed into the function
	r.GET("api/contract/:id", func(c *gin.Context) {

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
	})

	//TODO: Fix error handling of ids that don't exist
	//Return next 20 users after ID passed into the function
	r.GET("api/users/:lastID", func(c *gin.Context) {

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
	})

	//TODO: Fix error handling of ids that don't exist
	//Return next 20 assets after ID passed into the function
	r.GET("api/assets/:lastID", func(c *gin.Context) {

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
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
