package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	model *mongo.Collection
}

func NewUsers(i dependency) *Users {
	return &Users{
		model: i.Db.Collection("users"),
	}
}

func (x *Users) Get(ctx context.Context) (data []Users) {
	cursor, _ := x.model.Find(ctx, bson.M{})
	cursor.All(ctx, &data)
	return
}

func (x *Users) First() {

}

func (x *Users) Add() {

}

func (x *Users) Update() {

}

func (x *Users) Del() {

}
