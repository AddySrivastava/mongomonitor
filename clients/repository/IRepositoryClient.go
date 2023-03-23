package repository

import (
	"mongomonitor/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepositoryClient interface {
	InsertMany(docs []interface{}) types.WriteResponse
	InsertOne(docs interface{}) types.WriteResponse
	FindById() types.WriteResponse
	FindAll() types.WriteResponse
	Upsert(filter interface{}, update interface{}, options []byte) types.WriteResponse
	SetDatabase(database string)
	SetCollection(collection string)
	GetDatabase() string
	GetCollection() string
	ExecuteCommand(database string, cmd bson.D) map[string]interface{}
	ExecuteCommand2(database string, cmd bson.D) *mongo.SingleResult
}
