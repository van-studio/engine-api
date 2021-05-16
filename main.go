package main

import (
	"api/application"
	"api/bootstrap"
	curd "github.com/kainonly/gin-curd"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeDatabase,
			bootstrap.HttpServer,
			curd.Initialize,
		),
		fx.Invoke(application.Application),
	).Run()
}
