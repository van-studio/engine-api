package routes

import (
	"github.com/gin-gonic/gin"
	bit "github.com/kainonly/gin-bit"
	"github.com/kainonly/gin-helper/authx"
	"github.com/weplanx/api/controller"
)

func Initialize(
	r *gin.Engine,
	auth *authx.Auth,
	main *controller.Main,
	users *controller.Users,
	projects *controller.Projects,
) {
	r.GET("/", bit.Bind(main.Index))
	r.POST("/auth", bit.Bind(main.Login))
	r.GET("/auth", bit.Bind(main.Verify))
	r.DELETE("/auth", bit.Bind(main.Logout))

	middlewares := []gin.HandlerFunc{
		authx.Middleware(*auth),
	}
	rUsers := r.Group("/users", middlewares...)
	{
		rUsers.POST("/originLists", bit.Bind(users.OriginLists))
		rUsers.POST("/lists", bit.Bind(users.Lists))
		rUsers.POST("/get", bit.Bind(users.Get))
		rUsers.POST("/add", bit.Bind(users.Add))
		rUsers.POST("/edit", bit.Bind(users.Edit))
		rUsers.POST("/delete", bit.Bind(users.Delete))
	}
	rProjects := r.Group("projects", middlewares...)
	{
		rProjects.POST("/originLists", bit.Bind(projects.OriginLists))
		rProjects.POST("/lists", bit.Bind(projects.Lists))
		rProjects.POST("/get", bit.Bind(projects.Get))
		rProjects.POST("/add", bit.Bind(projects.Add))
		rProjects.POST("/edit", bit.Bind(projects.Edit))
		rProjects.POST("/delete", bit.Bind(projects.Delete))
	}

}
