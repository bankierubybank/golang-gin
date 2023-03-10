package route

import (
	"net/http"

	"github.com/bankierubybank/golang-gin/model"
	"github.com/gin-gonic/gin"
)

func Users(g *gin.RouterGroup) {
	g.GET("", GetUsers)
	g.GET(":id", GetUserByID)
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
func GetUsers(c *gin.Context) {
	us, err := model.GetUsers()
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
// @Accept		json
// @Param		id	path	int	true	"User ID"
// @Produce		json
// @Success		200
// @Router		/users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	u, err := model.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}
	c.JSON(http.StatusOK, u)
}
