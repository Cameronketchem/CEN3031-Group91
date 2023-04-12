package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/stretchr/testify/require"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Test Throttle
func TestRateLimiting(t *testing.T) {
	// Gin router with rate limiter middleware
	r := gin.Default()
	limiter := ratelimit.NewBucketWithQuantum(
		time.Second,
		10,
		10,
	)
	r.Use(RateLimiterMiddleware(limiter))

	// Endpoint handler function that just returns a simple message
	r.GET("/api/user/0", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Request handled"})
	})

	// Test server using the Gin router
	server := httptest.NewServer(r)
	defer server.Close()

	// Send 11 requests, bucket only should allow 10
	for i := 0; i < 11; i++ {
		resp, err := http.Get(server.URL + "/api/user/0")
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		if resp.StatusCode == http.StatusTooManyRequests {
			// Rate limit exceeded, passes test
			return
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Unexpected status code: %d", resp.StatusCode)
		}
	}

	// The rate limit was not exceeded so the test fails
	t.Fatalf("Rate limit was not exceeded")
}

// api/user/:id
func TestGetUserByID(t *testing.T) {

	api := New("db.sqlite", "secret", true)

	api.Store.Init(db.SynchronousModeFull, true)
	api.Store.Seed()
	defer api.Store.Close()

	r := setUpRouter()
	r.GET("/api/user/:id", func(c *gin.Context) {
		getUserById(&api.Store, c)
	})

	for i := 0; i < 7; i++ {
		path := "/api/user/" + strconv.Itoa(i)
		request, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, request)

		responseData, _ := ioutil.ReadAll(w.Body)

		user, err := api.Store.QueryUserById(i)
		require.NoError(t, err)

		expect, err := user.ToJSON()
		require.NoError(t, err)

		//Should show error if not the same
		require.Equal(t, expect, string(responseData))
	}
}

// api/asset/:id
func TestGetAssetByID(t *testing.T) {

	api := New("db.sqlite", "secret", true)

	api.Store.Init(db.SynchronousModeFull, true)
	api.Store.Seed()
	defer api.Store.Close()

	r := setUpRouter()
	r.GET("/api/asset/:id", func(c *gin.Context) {
		getAssetById(&api.Store, c)
	})

	for i := 0; i < 7; i++ {
		path := "/api/asset/" + strconv.Itoa(i)
		request, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, request)

		responseData, _ := ioutil.ReadAll(w.Body)

		asset, err := api.Store.QueryAssetById(i)
		require.NoError(t, err)

		//We need to format data to include objects within the assets
		EContract, err := api.Store.QueryContractById(asset.ExecContract)
		AContract, err := api.Store.QueryContractById(asset.AssetContract)
		Executor, _ := api.Store.QueryUserById(asset.Executor)
		require.NoError(t, err)

		JSONEcontract, err := EContract.ToJSON()
		JSONAcontract, err := AContract.ToJSON()
		JSONExecutor, err := Executor.ToJSON()
		JSONexpect := gin.H{"asset_id": asset.AssetID, "exec_contract": EContract, "asset_contract": AContract, "img_preview": asset.ImgPreview, "executor": Executor, "desc": asset.Desc, "price": asset.Price}

		jsonData, err := json.Marshal(JSONexpect)
		require.NoError(t, err)
		expect := string(jsonData)
		AContractContent := "\"asset_contract\":" + JSONAcontract
		EContractContent := "\"asset_contract\":" + JSONEcontract
		ExecutorContent := "\"executor\"" + JSONExecutor
		expect = strings.Replace(expect, ("\"asset_contract\":" + strconv.Itoa(asset.AssetContract)), AContractContent, 1)
		expect = strings.Replace(expect, ("\"exec_contract\":" + strconv.Itoa(asset.ExecContract)), EContractContent, 1)
		expect = strings.Replace(expect, ("\"executor\":" + strconv.Itoa(asset.Executor)), ExecutorContent, 1)

		//Should show error if not the same
		require.Equal(t, expect, string(responseData))
	}
}

// api/contract/:id
func TestGetContractByID(t *testing.T) {

	api := New("db.sqlite", "secret", true)

	api.Store.Init(db.SynchronousModeFull, true)
	api.Store.Seed()
	defer api.Store.Close()

	r := setUpRouter()
	r.GET("/api/contract/:id", func(c *gin.Context) {
		getContractById(&api.Store, c)
	})

	for i := 0; i < 7; i++ {
		path := "/api/contract/" + strconv.Itoa(i)
		request, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, request)

		responseData, _ := ioutil.ReadAll(w.Body)

		contract, err := api.Store.QueryContractById(i)
		require.NoError(t, err)

		expect, err := contract.ToJSON()
		require.NoError(t, err)

		//Should show error if not the same
		require.Equal(t, expect, string(responseData))
	}
}

// api/users/:lastID
func TestGetBatchUsers(t *testing.T) {
	api := New("db.sqlite", "secret", true)

	api.Store.Init(db.SynchronousModeFull, true)
	api.Store.Seed()
	defer api.Store.Close()

	r := setUpRouter()
	r.GET("/api/users/:lastID", func(c *gin.Context) {
		getBatchUsers(&api.Store, c)
	})

	path := "/api/users/" + strconv.Itoa(0)
	request, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	responseData, _ := ioutil.ReadAll(w.Body)

	users, err := api.Store.GetUsers(0)
	require.NoError(t, err)

	//Formats the expected output together correctly
	expect := "["
	line, err := users[0].ToJSON()
	expect = expect + line

	for i := 1; i < 6; i++ {
		line, err := users[i].ToJSON()
		require.NoError(t, err)

		expect = expect + "," + line
	}
	expect = expect + "]"

	//Should show error if not the same
	require.Equal(t, expect, string(responseData))

}

// api/users/:lastID
func TestGetBatchAssets(t *testing.T) {

	api := New("db.sqlite", "secret", true)

	api.Store.Init(db.SynchronousModeFull, true)
	api.Store.Seed()
	defer api.Store.Close()

	r := setUpRouter()
	r.GET("/api/assets/:lastID", func(c *gin.Context) {
		getBatchAssets(&api.Store, c)
	})

	path := "/api/assets/" + strconv.Itoa(0)
	request, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	responseData, _ := ioutil.ReadAll(w.Body)

	assets, err := api.Store.GetAssets(0)
	require.NoError(t, err)

	//Formats the expected output together correctly
	expect := "["

	//We need to format data to include objects within the assets
	EContract, err := api.Store.QueryContractById(assets[0].ExecContract)
	AContract, err := api.Store.QueryContractById(assets[0].AssetContract)
	Executor, _ := api.Store.QueryUserById(assets[0].Executor)
	require.NoError(t, err)

	JSONEcontract, err := EContract.ToJSON()
	JSONAcontract, err := AContract.ToJSON()
	JSONExecutor, err := Executor.ToJSON()
	JSONexpect := gin.H{"asset_id": assets[0].AssetID, "exec_contract": EContract, "asset_contract": AContract, "img_preview": assets[0].ImgPreview, "executor": Executor, "desc": assets[0].Desc, "price": assets[0].Price}

	jsonData, err := json.Marshal(JSONexpect)
	require.NoError(t, err)
	line := string(jsonData)
	AContractContent := "\"asset_contract\":" + JSONAcontract
	EContractContent := "\"asset_contract\":" + JSONEcontract
	ExecutorContent := "\"executor\"" + JSONExecutor
	line = strings.Replace(line, ("\"asset_contract\":" + strconv.Itoa(assets[0].AssetContract)), AContractContent, 1)
	line = strings.Replace(line, ("\"exec_contract\":" + strconv.Itoa(assets[0].ExecContract)), EContractContent, 1)
	line = strings.Replace(line, ("\"executor\":" + strconv.Itoa(assets[0].Executor)), ExecutorContent, 1)

	expect = expect + line

	for i := 1; i < 20; i++ {
		EContract, err := api.Store.QueryContractById(assets[i].ExecContract)
		AContract, err := api.Store.QueryContractById(assets[i].AssetContract)
		Executor, _ := api.Store.QueryUserById(assets[i].Executor)
		require.NoError(t, err)

		JSONEcontract, err := EContract.ToJSON()
		JSONAcontract, err := AContract.ToJSON()
		JSONExecutor, err := Executor.ToJSON()
		JSONexpect := gin.H{"asset_id": assets[i].AssetID, "exec_contract": EContract, "asset_contract": AContract, "img_preview": assets[i].ImgPreview, "executor": Executor, "desc": assets[i].Desc, "price": assets[i].Price}

		jsonData, err := json.Marshal(JSONexpect)
		line := string(jsonData)
		AContractContent := "\"asset_contract\":" + JSONAcontract
		EContractContent := "\"asset_contract\":" + JSONEcontract
		ExecutorContent := "\"executor\"" + JSONExecutor
		line = strings.Replace(line, ("\"asset_contract\":" + strconv.Itoa(assets[i].AssetContract)), AContractContent, 1)
		line = strings.Replace(line, ("\"exec_contract\":" + strconv.Itoa(assets[i].ExecContract)), EContractContent, 1)
		line = strings.Replace(line, ("\"executor\":" + strconv.Itoa(assets[i].Executor)), ExecutorContent, 1)

		expect = expect + "," + line
	}
	expect = expect + "]"

	//Should show error if not the same
	require.Equal(t, expect, string(responseData))
}
