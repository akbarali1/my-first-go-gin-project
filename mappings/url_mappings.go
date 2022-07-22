package mappings

import (
	"awesomeProject/Controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(Controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.GET("/users/:id", Controllers.GetUserDetail)
		v1.GET("/users/", Controllers.GetUser)
		v1.GET("/hello/", Controllers.Hello)
		v1.POST("/login/", Controllers.Login)
		v1.PUT("/users/:id", Controllers.UpdateUser)
		v1.POST("/users", Controllers.PostUser)
	}
}
