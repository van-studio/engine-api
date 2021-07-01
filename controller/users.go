package controller

import (
	bit "github.com/kainonly/gin-bit"
	"github.com/weplanx/api/model"
)

type Users struct {
	*bit.Crud
}

func NewUsers(b *bit.Bit) *Users {
	return &Users{
		Crud: b.Crud(model.User{}),
	}
}
