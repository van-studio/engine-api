package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kainonly/gin-helper/authx"
	"github.com/kainonly/gin-helper/hash"
	"github.com/weplanx/api/service"
	"gorm.io/gorm"
)

type Main struct {
	users *service.Users
	auth  *authx.Auth
}

func NewMain(users *service.Users, auth *authx.Auth) *Main {
	return &Main{users, auth}
}

func (x *Main) Index(c *gin.Context) interface{} {
	return gin.H{
		"msg": "hi",
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
	data, err := x.users.FindOne(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("email = ?", body.Email).
			Where("status = ?", true)
	})
	if err != nil {
		return err
	}
	if result, err := hash.Verify(body.Password, data.Password); err != nil || !result {
		return errors.New("the account does not exist or the password is incorrect")
	}
	if _, err := x.auth.Create(c, data.Email, data.ID, map[string]interface{}{}); err != nil {
		return err
	}
	return "ok"
}

func (x *Main) Verify(c *gin.Context) interface{} {
	if err := x.auth.Verify(c); err != nil {
		return err
	}
	return "ok"
}

func (x *Main) Logout(c *gin.Context) interface{} {
	if err := x.auth.Destory(c); err != nil {
		return err
	}
	return "ok"
}
