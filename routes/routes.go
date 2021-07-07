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
	projects *controller.Projects,
) {
	r.GET("/", bit.Bind(main.Index))
	r.POST("/login", bit.Bind(main.Login))
	r.GET("/verify", bit.Bind(main.Verify))
	r.DELETE("/logout", bit.Bind(main.Logout))

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
		rProjects.POST("/originLists", bit.Bind(projects.OriginLists))
		rProjects.POST("/lists", bit.Bind(projects.Lists))
		rProjects.POST("/get", bit.Bind(projects.Get))
		rProjects.POST("/add", bit.Bind(projects.Add))
		rProjects.POST("/edit", bit.Bind(projects.Edit))
		rProjects.POST("/delete", bit.Bind(projects.Delete))
	}

}
