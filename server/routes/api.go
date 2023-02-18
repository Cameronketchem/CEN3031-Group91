package routes

import (
	"fmt"
	"github.com/Cameronketchem/CEN3031-Group91/server/data"
	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
	"github.com/gin-gonic/gin"
)

type API struct {
	Store  data.Store
	Router *gin.Engine
	Ctx    *gin.Context
}

func New(dbpath string, inMemory bool) (api API) {
	api.Router = gin.Default()
	store, err := data.OpenStore(dbpath, inMemory)
	if err != nil {
		fmt.Errorf("Failed to open store, exiting with error: %s", err)
	}

	api.Store = store
	return api
}

func (api API) Start(port string, seed bool) {
	api.Store.Init(db.SynchronousModeFull, true)
	defer api.Store.Close()

	if seed {
		api.Store.Seed()
	}

	// server/routes/request.go
	api.get("api/user/:id", getUserById)
	api.get("api/asset/:id", getAssetById)
	api.get("api/contract/:id", getContractById)
	api.get("api/users/:lastID", getBatchUsers)
	api.get("api/assets/:lastID", getBatchAssets)

	api.Router.Run(port)
}

func (api API) get(route string, handler routeHandler) {
	api.Router.GET(route, func(c *gin.Context) {
		handler(&api.Store, c)
	})
}
