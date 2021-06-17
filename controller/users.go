package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/model"
	"github.com/weplanx/api/service"
	"time"
)

type Users struct {
	users *service.Users
}

func NewUsers(users *service.Users) *Users {
	return &Users{users}
}

func (x *Users) Lists(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := x.users.Get(ctx)

	c.JSON(200, gin.H{
		"message": "ok",
		"data":    data,
	})
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := x.users.FirstById(ctx, path.Id)

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := x.users.Insert(ctx, model.User{
		Email:      body.Email,
		Password:   body.Password,
		Name:       body.Name,
		Status:     true,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})

	c.JSON(201, gin.H{
		"message": "ok",
		"data": gin.H{
			"id": result.InsertedID,
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := x.users.UpdateById(ctx, path.Id, model.User{
		Name: body.Name,
	})

	c.JSON(200, gin.H{
		"message": "ok",
		"data": gin.H{
			"modified": result.ModifiedCount,
		},
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := x.users.Delete(ctx, path.Id)

	c.JSON(200, gin.H{
		"message": "ok",
		"data": gin.H{
			"deleted": result.DeletedCount,
		},
	})
}
