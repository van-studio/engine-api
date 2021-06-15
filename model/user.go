package model

type User struct {
	Email      string `bson:"email"`
	Password   string `bson:"password"`
	Name       string `bson:"name"`
	Status     bool   `bson:"status"`
	CreateTime uint64 `bson:"create_time"`
	UpdateTime uint64 `bson:"update_time"`
}
