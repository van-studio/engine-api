package model

import (
	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

type User struct {
	ID         uint   `json:"id"`
	Email      string `gorm:"type:varchar(255);unique" json:"email"`
	Password   string `gorm:"type:text" json:"-"`
	Name       string `gorm:"type:varchar(50)" json:"name"`
	Status     bool   `gorm:"default:1" json:"status"`
	CreateTime int64  `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime int64  `gorm:"autoUpdateTime" json:"update_time"`
}

func (x *User) BeforeSave(tx *gorm.DB) (err error) {
	if x.Password != "" {
		x.Password, _ = argon2id.CreateHash(x.Password, argon2id.DefaultParams)
	}
	return
}
