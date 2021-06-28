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
	data, err := x.users.FindMany()
	if err != nil {
		return err
	}
	return common.Ok{
		"data": data,
	}
}

func (x *Users) Get(c *gin.Context) interface{} {
	var path struct {
		Id string `uri:"id"`
	}
	if err := c.ShouldBindUri(&path); err != nil {
		return err
	}
	data, err := x.users.FindById(path.Id)
	if err != nil {
		return err
	}
	return common.Ok{
		"data": data,
	}
}

func (x *Users) Page(c *gin.Context) interface{} {
	var body common.ListByPage
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	data, err := x.users.Page(body)
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
		"result": result.RowsAffected,
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
		Status   bool   `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	result := x.users.UpdateById(path.Id, model.User{
		Password: body.Password,
		Name:     body.Name,
		Status:   &body.Status,
	})
	if result.Error != nil {
		return result.Error
	}
	return common.Ok{
		"result": result.RowsAffected,
	}
}

func (x *Users) Delete(c *gin.Context) interface{} {
	var path struct {
		Id string `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&path); err != nil {
		return err
	}
	result := x.users.Delete(path.Id)
	if result.Error != nil {
		return result.Error
	}
	return common.Ok{
		"result": result.RowsAffected,
	}
}
