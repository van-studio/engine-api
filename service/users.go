package service

import (
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
