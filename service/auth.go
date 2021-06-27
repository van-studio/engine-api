package service

import (
	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

type Auth struct {
	users *Users
}

func NewAuth(users *Users) *Auth {
	return &Auth{users: users}
}

func (x *Auth) Verify(email string, password string) (result bool, err error) {
	data, err := x.users.FindOne(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("email = ?", email).
			Where("status = ?", true)
	})
	if err != nil {
		return
	}
	result, err = argon2id.ComparePasswordAndHash(password, data.Password)
	if err != nil {
		return
	}
	return
}
