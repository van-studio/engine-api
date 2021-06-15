package controller

import (
	"github.com/weplanx/api/service"
	"go.uber.org/fx"
)

type dependency struct {
	fx.In

	Main *service.Main
}
