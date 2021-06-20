package model

type User struct {
	ID         uint   `json:"id"`
	Email      string `gorm:"type:varchar(20);unique" json:"email"`
	Password   string `gorm:"type:text" json:"-"`
	Name       string `gorm:"type:varchar(20)" json:"name"`
	Status     bool   `gorm:"default:1" json:"status"`
	CreateTime int64  `gorm:"default:0;autoCreateTime" json:"create_time"`
	UpdateTime int64  `gorm:"default:0;autoUpdateTime" json:"update_time"`
}
