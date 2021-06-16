package main

import (
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
			bootstrap.HttpServer,
			service.NewMain,
			service.NewUsers,
			controller.NewMain,
			controller.NewUsers,
		),
		fx.Invoke(routes.Initialize),
	).Run()
}
