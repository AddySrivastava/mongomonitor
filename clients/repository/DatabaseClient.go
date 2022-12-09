package repository

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	types "mongomonitor/types"

	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type DatabaseClient struct {
	ConnectionString string
	Options          map[string]string
	client           *mongo.Client
	database         string
	collection       string
}

func (DatabaseClient *DatabaseClient) InsertMany(docs []interface{}) types.WriteResponse {
	databaseClient := DatabaseClient.getDatabaseInstance()
	coll := databaseClient.Database(DatabaseClient.database).Collection(DatabaseClient.collection)

	_, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	return types.WriteResponse{Inserted: 0, Success: true}
}

func (DatabaseClient *DatabaseClient) Upsert(filter interface{}, update interface{}, options []byte) types.WriteResponse {

	databaseClient := DatabaseClient.getDatabaseInstance()

	coll := databaseClient.Database(DatabaseClient.database).Collection(DatabaseClient.collection)

	var opts mongoOptions.UpdateOptions

	bson.Unmarshal(options, &opts)

	_, err := coll.UpdateOne(context.TODO(), filter, update, &opts)

	if err != nil {
		fmt.Println(err)
	}

	return types.WriteResponse{Inserted: 0, Success: true}

}

func (DatabaseClient *DatabaseClient) InsertOne(docs interface{}) types.WriteResponse {
	databaseClient := DatabaseClient.getDatabaseInstance()
	coll := databaseClient.Database(DatabaseClient.database).Collection(DatabaseClient.collection)

	_, err := coll.InsertOne(context.TODO(), docs)

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

func (DatabaseClient *DatabaseClient) ExecuteCommand(cmd []byte) map[string]interface{} {

	var cmdResponse map[string]interface{}
	var bsonCommand bson.D

	err := bson.Unmarshal(cmd, &bsonCommand)

	if err != nil {
		panic(err)
	}

	databaseClient := DatabaseClient.getDatabaseInstance()
	db := databaseClient.Database(DatabaseClient.database)
	err2 := db.RunCommand(context.TODO(), bsonCommand).Decode(&cmdResponse)

	if err2 != nil {
		fmt.Println(err2)
	}

	return cmdResponse

}

func (DatabaseClient *DatabaseClient) SetDatabase(database string) {
	DatabaseClient.database = database
}

func (DatabaseClient *DatabaseClient) SetCollection(collection string) {
	DatabaseClient.collection = collection
}

func (DatabaseClient *DatabaseClient) getDatabaseInstance() *mongo.Client {

	if DatabaseClient.database == "" {
		DatabaseClient.SetDatabase("local")
	}
	if DatabaseClient.client != nil {
		return DatabaseClient.client
	}

	fmt.Println(DatabaseClient.ConnectionString)
	fmt.Println(DatabaseClient.client)

	clientOpts := mongoOptions.Client().ApplyURI(DatabaseClient.ConnectionString)
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
