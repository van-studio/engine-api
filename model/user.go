package model

import (
	"time"
)

type User struct {
	ID         uint64    `json:"_id"`
	Email      string    `gorm:"type:varchar(255);unique" json:"email" binding:"required,email"`
	Password   string    `gorm:"type:text" json:"-" binding:"required,min=12,max=20"`
	Name       string    `gorm:"type:varchar(50)" json:"name"`
	Status     *bool     `gorm:"default:true" json:"status"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
}
