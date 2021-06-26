package controller

import (
	"github.com/gin-gonic/gin"
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

	return gin.H{
		"message": "ok",
		"data":    data,
	}
}

func (x *Users) One(c *gin.Context) {
	var path struct {
		Id string `uri:"id"`
	}
	if err := c.ShouldBindUri(&path); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := x.users.FindOneById(path.Id)

	c.JSON(200, gin.H{
		"message": "ok",
		"data":    data,
	})
}

func (x *Users) Create(c *gin.Context) {
	var body struct {
		Email    string `binding:"required,email"`
		Password string `binding:"required,min=12,max=20"`
		Name     string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := model.User{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
	}
	if err := x.users.Create(data).Error; err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "ok",
		"data": gin.H{
			"id": data.ID,
		},
	})
}

func (x *Users) Update(c *gin.Context) {
	var path struct {
		Id string `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&path); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	var body struct {
		Password string `json:"password" binding:"omitempty,min=12,max=20"`
		Name     string `json:"name"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	x.users.UpdateById(path.Id, model.User{
		Name: body.Name,
	})

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (x *Users) Delete(c *gin.Context) {
	var path struct {
		Id string `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&path); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := x.users.Delete(path.Id).Error; err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}
