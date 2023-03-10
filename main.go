package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/bankierubybank/golang-gin/docs"
	"github.com/bankierubybank/golang-gin/route"
)

// @title			Golang Gin-Gonic Example API
// @version			v0.0.2
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
			debugRouter.GET("", GetDebug)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

type debugInfo struct {
	RuntimeInfo runtimeInfo `json:"runtimeInfo"`
	BuildInfo   buildInfo   `json:"buildInfo"`
}

type runtimeInfo struct {
	Hostname         string `json:"hostname"`
	UName            string `json:"uname"`
	GoRuntimeVersion string `json:"goruntimeversion"`
	K8Snode          string `json:"k8snode"`
	K8Snamespace     string `json:"k8snamespace"`
}
type buildInfo struct {
	GoBuildVersion string `json:"gobuildversion"`
	VCS            string `json:"vcs"`
	Commit         string `json:"commit"`
	CommitURL      string `json:"commiturl"`
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
func GetDebug(c *gin.Context) {
	d := new(debugInfo)

	d.RuntimeInfo.Hostname = os.Getenv("HOSTNAME")

	uname, unameErr := (exec.Command("uname", "-a")).Output()
	if unameErr == nil {
		d.RuntimeInfo.UName = strings.TrimRight(string(uname), "\n")
	}

	goruntimeversion, goruntimeversionErr := (exec.Command("go", "version")).Output()
	if goruntimeversionErr == nil {
		d.RuntimeInfo.GoRuntimeVersion = strings.TrimRight(string(goruntimeversion), "\n")
	}

	d.RuntimeInfo.K8Snode = os.Getenv("node")
	d.RuntimeInfo.K8Snamespace = os.Getenv("namespace")

	if info, ok := debug.ReadBuildInfo(); ok {
		d.BuildInfo.GoBuildVersion = info.GoVersion
		for _, setting := range info.Settings {
			if setting.Key == "vcs" {
				d.BuildInfo.VCS = setting.Value
			}
			if setting.Key == "vcs.revision" {
				d.BuildInfo.Commit = setting.Value
				d.BuildInfo.CommitURL = "https://github.com/bankierubybank/golang-gin/commit/" + setting.Value
			}
		}
	}

	c.JSON(http.StatusOK, d)
}
