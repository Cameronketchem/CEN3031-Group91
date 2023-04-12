package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Cameronketchem/CEN3031-Group91/server/data"
	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type API struct {
	Store  data.Store
	Router *gin.Engine
	Ctx    *gin.Context
	secret string
}

func RateLimiterMiddleware(limiter *ratelimit.Bucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.TakeAvailable(1) == 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func New(dbpath string, secret string, inMemory bool) (api API) {
	api.Router = gin.Default()

	//Allows 10 requests every 1 second
	limiter := ratelimit.NewBucketWithQuantum(
		time.Second,
		10,
		10,
	)

	api.Router.Use(RateLimiterMiddleware(limiter))

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
	// server/routes/auth.go
	api.get("api/user/:id", getUserById)
	api.get("api/asset/:id", getAssetById)
	api.get("api/contract/:id", getContractById)
	api.get("api/users/:lastID", getBatchUsers)
	api.get("api/assets/:lastID", getBatchAssets)
	api.get("api/auth/:addr", getNonce)
	api.post("api/auth/signup", postSignUp)
	api.post("api/auth/login", postLogIn)

	api.Router.Run(port)
}

func (api API) get(route string, handler routeHandler) {
	api.Router.GET(route, func(c *gin.Context) {
		// Enable CORS
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

		// Pass JWT secret
		c.Set("jwt-secret", api.secret)

		handler(&api.Store, c)
	})
}

func (api API) post(route string, handler routeHandler) {
	api.Router.POST(route, func(c *gin.Context) {
		// Enable CORS
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

		// Pass JWT secret
		c.Set("jwt-secret", api.secret)

		handler(&api.Store, c)
	})
}
