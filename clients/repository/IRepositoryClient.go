package repository

import "mongomonitor/types"

type IRepositoryClient interface {
	Insert(docs []interface{}) types.WriteResponse
	FindById() types.WriteResponse
	FindAll() types.WriteResponse
	SetDatabase(database string)
	SetCollection(collection string)
}
