package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/controller"
)

func Initialize(
	route *gin.Engine,
	main *controller.Main,
) {
	route.GET("/", main.Index)
}
