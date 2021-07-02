package service

import (
	"github.com/kainonly/gin-helper/hash"
	"github.com/weplanx/api/model"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) *Users {
	return &Users{
		db: db,
	}
}

func (x *Users) FindOne(query Query) (data model.User, err error) {
	if err = query(x.db).First(&data).Error; err != nil {
		return
	}
	return
}

func (x *Users) Data(user model.User) model.User {
	var password string
	if user.Password != "" {
		password, _ = hash.Make(user.Password)
	}
	return model.User{
		Email:    user.Email,
		Password: password,
		Name:     user.Name,
		Status:   user.Status,
	}
}
