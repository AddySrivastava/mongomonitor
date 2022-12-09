package repository

import "mongomonitor/types"

type IRepositoryClient interface {
	InsertMany(docs []interface{}) types.WriteResponse
	InsertOne(docs interface{}) types.WriteResponse
	FindById() types.WriteResponse
	FindAll() types.WriteResponse
	Upsert(filter interface{}, update interface{}, options []byte) types.WriteResponse
	SetDatabase(database string)
	SetCollection(collection string)
	ExecuteCommand(cmd []byte) map[string]interface{}
}
