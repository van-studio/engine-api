package main

import (
	"github.com/weplanx/api/bootstrap"
	"github.com/weplanx/api/model"
)

func main() {
	cfg := bootstrap.LoadConfiguration()
	db := bootstrap.InitializeDatabase(cfg)
	db = db.Debug()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Activities{})
	db.AutoMigrate(&model.Logs{})
	db.AutoMigrate(&model.Project{})
}
