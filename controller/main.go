package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Main struct {
	*dependency
}

func NewMain(i dependency) *Main {
	return &Main{&i}
}

func (m *Main) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": m.Main.Index(),
	})
	return
}

func (m *Main) Login(c *gin.Context) {
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
