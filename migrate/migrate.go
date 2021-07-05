package main

import (
	"github.com/weplanx/api/bootstrap"
	"github.com/weplanx/api/model"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db = db.Debug()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Project{})
}

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeDatabase,
		),
		fx.Invoke(migrate),
	).Run()
}
