package routes

import (
	"github.com/gin-gonic/gin"
	bit "github.com/kainonly/gin-bit"
	"github.com/weplanx/api/controller"
)

func Initialize(
	r *gin.Engine,
	main *controller.Main,
	users *controller.Users,
) {
	r.GET("/", bit.Bind(main.Index))
	rUsers := r.Group("/users")
	{
		rUsers.POST("/originLists", bit.Bind(users.OriginLists))
		rUsers.POST("/lists", bit.Bind(users.Lists))
		rUsers.POST("/get", bit.Bind(users.Get))
		rUsers.POST("/add", bit.Bind(users.Add))
		rUsers.POST("/edit", bit.Bind(users.Edit))
		rUsers.POST("/delete", bit.Bind(users.Delete))
	}
	rProjects := r.Group("projects")
	{
		rProjects.POST("/originLists", bit.Bind(users.OriginLists))
		rProjects.POST("/lists", bit.Bind(users.Lists))
		rProjects.POST("/get", bit.Bind(users.Get))
		rProjects.POST("/add", bit.Bind(users.Add))
		rProjects.POST("/edit", bit.Bind(users.Edit))
		rProjects.POST("/delete", bit.Bind(users.Delete))
	}

}
