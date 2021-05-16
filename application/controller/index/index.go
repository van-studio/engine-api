package index

import (
	"api/application/common"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	common.Dependency
}

type LoginBody struct {
	Username string `binding:"required,min=4,max=20"`
	Password string `binding:"required,min=12,max=20"`
}

func (c *Controller) Login(ctx *gin.Context) interface{} {
	return true
}
