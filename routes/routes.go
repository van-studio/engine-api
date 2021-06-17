package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/controller"
)

func Initialize(
	route *gin.Engine,
	main *controller.Main,
	users *controller.Users,
) {
	route.GET("/", main.Index)
	route.POST("/login", main.Login)

	route.GET("/users", users.Lists)
	route.POST("/users", users.Create)
	route.DELETE("/users/:id", users.Delete)

}
