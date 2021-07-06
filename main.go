package main

import (
	bit "github.com/kainonly/gin-bit"
	"github.com/weplanx/api/bootstrap"
	"github.com/weplanx/api/controller"
	"github.com/weplanx/api/routes"
	"github.com/weplanx/api/service"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeDatabase,
			bootstrap.InitializeAuth,
			bootstrap.HttpServer,
			bit.Initialize,
		),
		service.Provides,
		controller.Provides,
		fx.Invoke(routes.Initialize),
	).Run()
}
