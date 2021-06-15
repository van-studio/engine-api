package controller

import (
	"github.com/gin-gonic/gin"
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
