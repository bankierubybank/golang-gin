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

func TestGetAlbums(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums", getAlbums)
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var albums []album
	json.Unmarshal(w.Body.Bytes(), &albums)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDebug(t *testing.T) {
	r := SetUpRouter()
	r.GET("/debug", debug)
	req, _ := http.NewRequest("GET", "/debug", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var d debugInfo
	json.Unmarshal(w.Body.Bytes(), &d)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println("W.BODY ", w.Body)
}
