package main

import (
	"github.com/weplanx/api/bootstrap"
	"github.com/weplanx/api/routes"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeDatabase,
			bootstrap.HttpServer,
		),
		fx.Invoke(routes.Initialize),
	).Run()
}
