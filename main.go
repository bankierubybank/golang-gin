package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/bankierubybank/golang-gin/docs"
	"github.com/bankierubybank/golang-gin/route"
)

// @title			Golang Gin-Gonic Swagger Example API
// @version			v0.0.2
// @license.name	Apache 2.0
func main() {
	router := gin.Default()

	// CORS for http://localhost:5173 origin, allowing:
	// - GET, PUT, and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		route.Users(v1.Group("/users"))
		debugRouter := v1.Group("/debug")
		{
			debugRouter.GET("", debug)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

type debugInfo struct {
	Hostname  string `json:"hostname"`
	UName     string `json:"uname"`
	GoVersion string `json:"goversion"`
}

// @BasePath	/api/v1
// @Summary		Get debug information
// @Schemes
// @Description	Get debug information
// @Tags		debug
// @Accept		json
// @Produce		json
// @Success		200
// @Router		/debug [get]
func debug(c *gin.Context) {
	d := new(debugInfo)

	hostname, hostnameErr := (exec.Command("hostname")).Output()
	var h string = strings.TrimRight(string(hostname), "\n")
	if hostnameErr != nil {
		d.Hostname = ""
	} else {
		d.Hostname = h
	}

	uname, unameErr := (exec.Command("uname", "-a")).Output()
	var u string = strings.TrimRight(string(uname), "\n")
	if unameErr != nil {
		d.UName = ""
	} else {
		d.UName = u
	}

	goversion, goversionErr := (exec.Command("go", "version")).Output()
	var g string = strings.TrimRight(string(goversion), "\n")
	if goversionErr != nil {
		d.GoVersion = ""
	} else {
		d.GoVersion = g
	}

	c.IndentedJSON(http.StatusOK, d)
}
