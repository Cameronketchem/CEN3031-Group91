package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// api/user/:id
func TestGetUserByID(t *testing.T) {

	api := New("fn.db", true)

	Request, err := http.NewRequest("GET", "api/user", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}

	Recorder := httptest.NewRecorder()

	api.Router.ServeHTTP(Recorder, Request)

	/*if Recorder.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", Recorder.Code, " instead of ", http.StatusOK)
	}*/

	assert.Equal(t, http.StatusOK, Recorder.Code)
}

// api/asset/:id
func TestGetAssetByID(t *testing.T) {
	_, err := http.NewRequest("GET", "api/asset/", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
}

// api/contract/:id
func TestGetContractByID(t *testing.T) {
	_, err := http.NewRequest("GET", "api/contract/", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
}

// api/users/:lastID
func TestGetBatchUsers(t *testing.T) {
	_, err := http.NewRequest("GET", "api/users/", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
}

// api/users/:lastID
func TestGetBatchAssets(t *testing.T) {
	_, err := http.NewRequest("GET", "api/assets/", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
}
