package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weplanx/api/service"
	"log"
)

type Main struct {
	main *service.Main
}

func NewMain(main *service.Main) *Main {
	return &Main{main}
}

func (x *Main) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": x.main.Index(),
	})
	return
}

func (x *Main) Login(c *gin.Context) {
	var body struct {
		Email    string `binding:"required,email"`
		Password string `binding:"required,min=12,max=20"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(body)
}
