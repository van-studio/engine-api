package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/common"
	"github.com/weplanx/api/model"
	"github.com/weplanx/api/service"
)

type Users struct {
	users *service.Users
}

func NewUsers(users *service.Users) *Users {
	return &Users{users}
}

func (x *Users) Lists(c *gin.Context) interface{} {
	data, err := x.users.Find()
	if err != nil {
		return err
	}
	return common.Ok{
		"data": data,
	}
}

func (x *Users) One(c *gin.Context) interface{} {
	var path struct {
		Id string `uri:"id"`
	}
	if err := c.ShouldBindUri(&path); err != nil {
		return err
	}
	data, err := x.users.FindOneById(path.Id)
	if err != nil {
		return err
	}
	return common.Ok{
		"data": data,
	}
}

func (x *Users) Create(c *gin.Context) interface{} {
	var body struct {
		Email    string `binding:"required,email"`
		Password string `binding:"required,min=12,max=20"`
		Name     string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	data := model.User{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
	}
	result := x.users.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return common.Create{
		"affected": result.RowsAffected,
	}
}

func (x *Users) Update(c *gin.Context) interface{} {
	var path struct {
		Id string `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&path); err != nil {
		return err
	}
	var body struct {
		Password string `json:"password" binding:"omitempty,min=12,max=20"`
		Name     string `json:"name"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	if err := x.users.UpdateById(path.Id, model.User{
		Name: body.Name,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (x *Users) Delete(c *gin.Context) interface{} {
	var path struct {
		Id string `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&path); err != nil {
		return err
	}
	if err := x.users.Delete(path.Id).Error; err != nil {
		return err
	}
	return nil
}
