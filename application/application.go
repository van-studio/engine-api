package application

import (
	"api/application/common"
	"api/application/controller/index"
	"api/routes"
	"github.com/gin-gonic/gin"
	"github.com/kainonly/gin-extra/mvcx"
)

func Application(router *gin.Engine, dependency common.Dependency) {
	router.GET("/", routes.Default)
	system := router.Group("/system")
	{
		//unifyMiddleware := []mvcx.Middleware{}
		mvc := mvcx.Initialize(system, dependency)
		mvc.AutoController("/main", &index.Controller{}, mvcx.Middleware{})
	}
}
