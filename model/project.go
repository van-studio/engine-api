package model

import "time"

type Project struct {
	ID          uint      `json:"_id"`
	Key         string    `gorm:"varchar(20);unique" json:"key"`
	Name        string    `gorm:"type:varchar(50)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Status      *bool     `gorm:"default:true" json:"status"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"update_time"`
}
