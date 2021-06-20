package service

import (
	"context"
	"github.com/weplanx/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Users struct {
	model *mongo.Collection
}

func NewUsers(Db *mongo.Database) *Users {
	return &Users{
		model: Db.Collection("users"),
	}
}

func (x *Users) Get(ctx context.Context) (data []model.User) {
	opts := options.Find()
	opts.Projection = map[string]interface{}{
		"password": false,
	}
	cursor, _ := x.model.Find(ctx, bson.M{}, opts)
	cursor.All(ctx, &data)
	return
}

func (x *Users) Page(ctx context.Context) {

}

func (x *Users) First(ctx context.Context, filter interface{}) (data model.User) {
	opts := options.FindOne()
	opts.Projection = map[string]interface{}{
		"password": false,
	}
	x.model.FindOne(ctx, filter, opts).Decode(&data)
	return
}

func (x *Users) FirstById(ctx context.Context, id string) model.User {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return x.First(ctx, bson.M{"_id": objectId})
}

func (x *Users) Insert(ctx context.Context, data model.User) (result *mongo.InsertOneResult) {
	data.CreateTime = time.Now()
	data.UpdateTime = time.Now()
	result, _ = x.model.InsertOne(ctx, data)
	return
}

func (x *Users) Update(ctx context.Context, filter interface{}, data model.User) (result *mongo.UpdateResult) {
	data.UpdateTime = time.Now()
	result, _ = x.model.ReplaceOne(ctx, filter, data)
	return
}

func (x *Users) UpdateById(ctx context.Context, id string, data model.User) *mongo.UpdateResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return x.Update(ctx, bson.M{"_id": objectId}, data)
}

func (x *Users) Delete(ctx context.Context, id string) (result *mongo.DeleteResult) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, _ = x.model.DeleteOne(ctx, bson.M{"_id": objectId})
	return
}
