package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"

	"github.com/bankierubybank/golang-gin/model"
	"github.com/bankierubybank/golang-gin/route"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetUsers(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/users", route.GetUsers)
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var users []model.UserModel
	json.Unmarshal(w.Body.Bytes(), &users)

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
