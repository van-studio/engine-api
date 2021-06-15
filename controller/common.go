package controller

import (
	"github.com/weplanx/api/service"
	"go.uber.org/fx"
)

type services struct {
	fx.In

	Index *service.IndexService
}
