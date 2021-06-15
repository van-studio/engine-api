package service

import (
	"github.com/weplanx/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

type dependency struct {
	fx.In

	Config *config.Config
	Db     *mongo.Database
}
