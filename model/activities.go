package model

import "time"

type Activities struct {
	ID     uint64    `json:"_id"`
	UserID uint64    `json:"user_id"`
	Ip     string    `gorm:"type:varchar(128)" json:"ip"`
	Device string    `gorm:"type:text" json:"device"`
	Time   time.Time `gorm:"autoCreateTime;index" json:"time"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
