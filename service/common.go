package service

import "go.uber.org/fx"

var Provides = fx.Provide(
	NewMain,
	NewUsers,
)
