package controller

import (
	"github.com/gin-gonic/gin"
	bit "github.com/kainonly/gin-bit"
	"github.com/weplanx/api/model"
	"github.com/weplanx/api/service"
)

type Users struct {
	*bit.Crud
	users *service.Users
}

func NewUsers(b *bit.Bit, users *service.Users) *Users {
	return &Users{
		Crud:  b.Crud(model.User{}),
		users: users,
	}
}

func (x *Users) Add(c *gin.Context) interface{} {
	var body struct {
		Email    string `binding:"required,email"`
		Password string `binding:"required,min=12,max=20"`
		Name     string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	data := x.users.Data(model.User{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
	})
	bit.Mixed(c, bit.SetData(&data))
	return x.Crud.Add(c)
}

func (x *Users) Edit(c *gin.Context) interface{} {
	var body struct {
		bit.EditBody
		Email    string `binding:"omitempty,email"`
		Password string `binding:"omitempty,min=12,max=20"`
		Name     string
		Status   *bool
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	data := x.users.Data(model.User{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
		Status:   body.Status,
	})
	bit.Mixed(c,
		bit.SetBody(&body),
		bit.SetData(&data),
	)
	return x.Crud.Edit(c)
}
