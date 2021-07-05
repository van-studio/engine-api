package controller

import (
	bit "github.com/kainonly/gin-bit"
	"github.com/weplanx/api/model"
)

type Projects struct {
	*bit.Crud
}

func NewProjects(b *bit.Bit) *Projects {
	return &Projects{
		Crud: b.Crud(model.Project{}),
	}
}
