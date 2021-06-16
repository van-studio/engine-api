package controller

import (
	"context"
	"github.com/gin-gonic/gin"
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

	c.JSON(200, data)
}
