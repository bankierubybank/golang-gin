package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/bankierubybank/golang-gin/docs"
	_ "github.com/bankierubybank/golang-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Swagger Example API
// @version			v0.0.1
// @license.name	Apache 2.0
func main() {
	// setup mock data for api

	router := gin.Default()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
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
		users := v1.Group("/users")
		{
			users.GET("", get_users)
			users.GET(":id", get_user_id)
		}
		random := v1.Group("/cat")
		{
			random.GET("/random", getRandomCat)
		}
		debugRouter := v1.Group("/debug")
		{
			debugRouter.GET("", debug)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

// @BasePath	/api/v1
// @Summary		Get all users
// @Schemes
// @Description	Get all users
// @Tags		users
// @Accept		json
// @Produce		json
// @Success		200
// @Router		/users/ [get]
func get_users(c *gin.Context) {
	us, err := model.getUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users"})
	}
	c.JSON(http.StatusOK, us)
}

// @BasePath	/api/v1
// @Summary		Get an user by ID
// @Schemes
// @Description	Get an user by ID
// @Tags		users
// @Accept		jsogn
// @Param		id	path	int	true	"User ID"
// @Produce		json
// @Success		200
// @Router		/users/{id} [get]
func get_user_id(c *gin.Context) {
	id := c.Param("id")

	u, err := model.getUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}
	c.JSON(http.StatusOK, u)
}

// @BasePath	/api/v1
// @Summary		Get random cat
// @Schemes
// @Description	Get random cat
// @Tags		cat
// @Accept		json
// @Produce		json
// @Success		200
// @Router		/cat/random [get]
func getRandomCat(c *gin.Context) {
	var code [5]int
	code[0] = 200
	code[1] = 200
	code[2] = 403
	code[3] = 404
	code[4] = 503
	min := 1
	max := 5
	var index = rand.Intn(max-min) + min
	var get = "https://http.cat/" + string(rune(code[index]))
	resp, err := http.Get(get)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, "")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, "")
	}

	c.IndentedJSON(http.StatusOK, body)
}

type debugInfo struct {
	Hostname  string `json:"hostname"`
	UName     string `json:"uname"`
	GoVersion string `json:"goversion"`
	ClientIP  string `json:"clientip"`
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
	if c.ClientIP() != "" {
		d.ClientIP = c.ClientIP()
	} else {
		d.ClientIP = ""
	}
	fmt.Println("IP: ", c.ClientIP())

	hostname, hostnameErr := (exec.Command("hostname")).Output()
	var h string = strings.TrimRight(string(hostname), "\n")
	if hostnameErr != nil {
		d.Hostname = ""
	} else {
		d.Hostname = h
	}
	fmt.Println("H: ", h)

	uname, unameErr := (exec.Command("uname", "-a")).Output()
	var u string = strings.TrimRight(string(uname), "\n")
	if unameErr != nil {
		d.UName = ""
	} else {
		d.UName = u
	}
	fmt.Println("U: ", u)

	goversion, goversionErr := (exec.Command("go", "version")).Output()
	var g string = strings.TrimRight(string(goversion), "\n")
	if goversionErr != nil {
		d.GoVersion = ""
	} else {
		d.GoVersion = g
	}
	fmt.Println("G: ", g)

	c.IndentedJSON(http.StatusOK, d)
}
