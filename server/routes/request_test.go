package routes

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// api/user/:id
func TestGetUserByID(t *testing.T) {

	api := New("db.sqlite", true)

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

	api := New("db.sqlite", true)

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

		expect, err := asset.ToJSON()
		require.NoError(t, err)

		//Should show error if not the same
		require.Equal(t, expect, string(responseData))
	}
}

// api/contract/:id
func TestGetContractByID(t *testing.T) {

	api := New("db.sqlite", true)

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
	api := New("db.sqlite", true)

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

	api := New("db.sqlite", true)

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
	line, err := assets[0].ToJSON()
	expect = expect + line

	for i := 1; i < 20; i++ {
		line, err := assets[i].ToJSON()
		require.NoError(t, err)

		expect = expect + "," + line
	}
	expect = expect + "]"

	//Should show error if not the same
	require.Equal(t, expect, string(responseData))
}
