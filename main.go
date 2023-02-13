package main

import (
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"

	"fmt"

	docs "github.com/bankierubybank/golang-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @title			Swagger Example API
// @version			v0.0.1
// @license.name	Apache 2.0
func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			eg.GET("/albums", getAlbums)
			eg.GET("/albums/:id", getAlbumByID)
			eg.POST("/albums", postAlbums)
			eg.GET("/debug", debug)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

// @BasePath	/api/v1
// @Summary		Get all albums
// @Schemes
// @Description	Get all albums
// @Tags		albums
// @Accept		json
// @Produce		json
// @Success		200
// @Router		/albums/ [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// @BasePath	/api/v1
// @Summary		Create an album
// @Schemes
// @Description	Create an album
// @Tags		albums
// @Accept		json
// @Produce		json
// @Success		200
// @Router		/albums/ [post]
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// @BasePath	/api/v1
// @Summary		Get an album by ID
// @Schemes
// @Description	Get an album by ID
// @Tags		albums
// @Accept		json
// @Param		id	path	int	true	"Album ID"
// @Produce		json
// @Success		200
// @Router		/albums/{id} [get]
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
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
