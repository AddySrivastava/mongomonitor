package repository

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	types "mongomonitor/types"

	mongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseClient struct {
	ConnectionString string
	Options          map[string]string
	client           *mongo.Client
	database         string
	collection       string
}

func (DatabaseClient *DatabaseClient) Insert(docs []interface{}) types.WriteResponse {
	databaseClient := DatabaseClient.getDatabaseInstance()
	coll := databaseClient.Database(DatabaseClient.database).Collection(DatabaseClient.collection)

	_, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	return types.WriteResponse{Inserted: 0, Success: true}
}

func (DatabaseClient *DatabaseClient) FindById() types.WriteResponse {
	return types.WriteResponse{Inserted: 0, Success: true}
}

func (DatabaseClient *DatabaseClient) FindAll() types.WriteResponse {
	return types.WriteResponse{Inserted: 0, Success: true}
}
func (DatabaseClient *DatabaseClient) SetDatabase(database string) {
	DatabaseClient.database = database
}

func (DatabaseClient *DatabaseClient) SetCollection(collection string) {
	DatabaseClient.collection = collection
}

func (DatabaseClient *DatabaseClient) getDatabaseInstance() *mongo.Client {

	clientOpts := options.Client().ApplyURI(DatabaseClient.ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	response, _ := json.Marshal(client)

	fmt.Println(string(response))

	DatabaseClient.client = client

	if err != nil {
		panic(err)
	} else {
		return client
	}
}
