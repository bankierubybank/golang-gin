package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetUsers(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/users", getUsers)
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var users []user
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRandomCat(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/cat/random", getRandomCat)
	req, _ := http.NewRequest("GET", "/api/v1/cat/random", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDebug(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/debug", debug)
	req, _ := http.NewRequest("GET", "/api/v1/debug", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var d debugInfo
	json.Unmarshal(w.Body.Bytes(), &d)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println("W.BODY ", w.Body)
}
