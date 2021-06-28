package service

import (
	"github.com/alexedwards/argon2id"
	"github.com/weplanx/api/common"
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

func (x *Users) FindMany() (data []model.User, err error) {
	err = x.db.Omit("password").Find(&data).Error
	return
}

func (x *Users) FindOne(query Query) (data model.User, err error) {
	err = query(x.db).First(&data).Error
	return
}

func (x *Users) FindById(id interface{}) (data model.User, err error) {
	err = x.db.Omit("password").First(&data, id).Error
	return
}

func (x *Users) Page(p common.ListByPage) (data []model.User, err error) {
	page := p.Pagination
	err = x.db.
		Limit(page.Limit).
		Offset((page.Index - 1) * page.Limit).
		Find(&data).Error
	return
}

func (x *Users) Create(data model.User) *gorm.DB {
	if data.Password != "" {
		data.Password, _ = argon2id.CreateHash(data.Password, argon2id.DefaultParams)
	}
	return x.db.Create(&data)
}

func (x *Users) Update(query Query, data model.User) *gorm.DB {
	return query(x.db).Updates(data)
}

func (x *Users) UpdateById(id interface{}, data model.User) *gorm.DB {
	if data.Password != "" {
		data.Password, _ = argon2id.CreateHash(data.Password, argon2id.DefaultParams)
	}
	return x.db.Model(&model.User{}).Where("id = ?", id).Updates(data)
}

func (x *Users) Delete(id interface{}) *gorm.DB {
	return x.db.Delete(&model.User{}, id)
}
