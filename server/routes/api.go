package routes

import (
	"fmt"
	"github.com/Cameronketchem/CEN3031-Group91/server/blockchain"
	"github.com/Cameronketchem/CEN3031-Group91/server/data"
	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

// API uses gin to host an instance of the API. Must be created using New
// and started with api.Start().
type API struct {
	Store    data.Store
	Router   *gin.Engine
	Ctx      *gin.Context
	Contract blockchain.Executor

	// Config data.
	secret string
}

// Enforces a ceiling to the number of allowable API requests per second.
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

// Creates and initializes a new API instance.
func New(dbpath, secret, privKey, nodeAddr string, inMemory bool) (api API) {
	api.Router = gin.Default()
	api.secret = secret

	// Set rate limiter middleware.
	limiter := ratelimit.NewBucketWithQuantum(time.Second, 10, 10)

	api.Router.Use(
		RateLimiterMiddleware(limiter),
		func(c *gin.Context) {

			//Allows site administators to control resources allowed to be loaded for a page.
			//Helps protect from XSS attacks. Default src-directve set for the fetch directives of each resource.
			c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")

			//Avoid click-jacking attacks by only including a page in <frame> that is served by the same place as the page.
			c.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")

			//Browser remembers site should only be accessed by HTTPS for 1 year.
			c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

			c.Next()
		})

	// Open datastore
	store, err := data.OpenStore(dbpath, inMemory)
	if err != nil {
		fmt.Errorf("Failed to open store, exiting with error: %s", err)
	}
	api.Store = store

	// Create blockchain contract executor.
	contract := blockchain.NewExecutor(privKey, nodeAddr)
	api.Contract = contract

	return api
}

// Starts the API - should be internet addressible.
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

	// Smart-contract dependent routes.
	api.get("api/asset/:id/price", getAssetPrice)
	api.get("api/asset/:id/active", getAssetActive)
	api.get("api/asset/:id/contrib/:addr", getAssetContribByAddr)

	// TODO: balance is not pooledAmt
	api.get("api/asset/:id/balance", getAssetBalance)

	// Authenticated routes.
	api.post("api/asset", isLoggedIn, postAsset)

	// Run gin API process.
	api.Router.Run(port)
}

// --- GIN WRAPPERS ---
// Wraps gin request functions, that way headers, datastore, and secrets
// can be passed via context.

func (api API) get(route string, handlers ...routeHandler) {
	api.Router.GET(route, func(c *gin.Context) {
		if len(handlers) < 1 {
			panic("Expected at least 1 middleware handler")
		}

		// Enable CORS
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

		// Pass JWT secret
		c.Set("jwt-secret", api.secret)

		// Run Handlers
		for i := 0; i < len(handlers) && !c.IsAborted(); i++ {
			handlers[i](&api.Store, &api.Contract, c)
		}
	})
}

func (api API) post(route string, handlers ...routeHandler) {
	api.Router.POST(route, func(c *gin.Context) {
		if len(handlers) < 1 {
			panic("Expected at least 1 middleware handler")
		}

		// Enable CORS
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

		// Pass JWT secret
		c.Set("jwt-secret", api.secret)

		// Run Handlers
		for i := 0; i < len(handlers) && !c.IsAborted(); i++ {
			handlers[i](&api.Store, &api.Contract, c)
		}
	})
}
