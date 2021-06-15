package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/controller"
)

func Initialize(
	router *gin.Engine,
	index *controller.IndexController,
) {
	router.GET("/", index.Index)
}
