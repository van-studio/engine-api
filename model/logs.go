package model

import "time"

type Logs struct {
	ID     uint64    `json:"_id"`
	UserID uint64    `json:"user_id"`
	Topic  string    `gorm:"type:varchar(200)" json:"topic"`
	Action string    `gorm:"type:text" json:"action"`
	Time   time.Time `gorm:"autoCreateTime;index" json:"time"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
