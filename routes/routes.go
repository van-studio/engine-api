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
	routes := [][]interface{}{
		{"GET", "/", main.Index},
		{"POST", "/login", main.Login},

		{"GET", "/users", users.Lists},
		{"GET", "/users/:id", users.One},
		{"POST", "/users", users.Create},
		{"PUT", "/users/:id", users.Update},
		{"DELETE", "/users/:id", users.Delete},
	}

	for _, r := range routes {
		handlers := []gin.HandlerFunc{bind(r[2])}
		for _, ext := range r[3:] {
			handlers = append(handlers, ext.(gin.HandlerFunc))
		}
		route.Handle(r[0].(string), r[1].(string), handlers...)
	}
}

func bind(handlerFn interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if fn, ok := handlerFn.(func(ctx *gin.Context) interface{}); ok {
			switch result := fn(ctx).(type) {
			case gin.H:
				ctx.JSON(200, result)
				break
			case error:
				ctx.JSON(400, gin.H{
					"msg": result.Error(),
				})
				break
			default:
				ctx.JSON(200, gin.H{
					"msg": "ok",
				})
			}
		}
	}
}
