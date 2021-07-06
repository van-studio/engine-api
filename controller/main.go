package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/service"
)

type Main struct {
	auth *service.Auth
}

func NewMain(auth *service.Auth) *Main {
	return &Main{auth}
}

func (x *Main) Index(c *gin.Context) interface{} {
	return gin.H{
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
	if err := x.auth.Verify(body.Email, body.Password); err != nil {
		return err
	}
	return "ok"
}
