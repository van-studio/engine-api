package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/common"
	"github.com/weplanx/api/service"
)

type Main struct {
	main *service.Main
}

func NewMain(main *service.Main) *Main {
	return &Main{main}
}

func (x *Main) Index(c *gin.Context) interface{} {
	return common.Ok{
		"msg": x.main.Index(),
	}
}

func (x *Main) Login(c *gin.Context) interface{} {
	var body struct {
		Email    string `binding:"required,email"`
		Password string `binding:"required,min=12,max=20"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	return nil
}
