package service

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Provides = fx.Provide(
	NewAuth,
	NewUsers,
)

type Query func(tx *gorm.DB) *gorm.DB
type Result struct {
	Error        error
	RowsAffected int64
}
