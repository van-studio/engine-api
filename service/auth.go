package service

import (
	"errors"
	"github.com/kainonly/gin-helper/hash"
	"gorm.io/gorm"
)

type Auth struct {
	users *Users
}

func NewAuth(users *Users) *Auth {
	return &Auth{users: users}
}

func (x *Auth) Verify(email string, password string) error {
	data, err := x.users.FindOne(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("email = ?", email).
			Where("status = ?", true)
	})
	if err != nil {
		return err
	}
	result, err := hash.Verify(password, data.Password)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("the account does not exist or the password is incorrect")
	}
	return nil
}
