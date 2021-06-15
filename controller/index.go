package controller

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	*services
}

func NewIndex(services services) *IndexController {
	return &IndexController{&services}
}

func (m *IndexController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": m.services.Index.Index(),
	})
	return
}
