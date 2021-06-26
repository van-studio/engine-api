package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/common"
	"github.com/weplanx/api/service"
)

type Main struct {
	auth *service.Auth
}

func NewMain(auth *service.Auth) *Main {
	return &Main{auth}
}

func (x *Main) Index(c *gin.Context) interface{} {
	return common.Ok{
		"msg": "hello",
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
	result, err := x.auth.Verify(body.Email, body.Password)
	if err != nil {
		return err
	}
	if result {
		// TODO: PASSPORT
	}
	return nil
}
