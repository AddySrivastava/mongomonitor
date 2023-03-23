package repository

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	types "mongomonitor/types"
	"mongomonitor/utils"

	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

//init logger

var logger = utils.GetLogger()

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

func (DatabaseClient *DatabaseClient) ExecuteCommand2(database string, cmd bson.D) *mongo.SingleResult {

	//var bsonCommand bson.D

	//bson.Unmarshal(cmd, &bsonCommand)

	fmt.Printf("%v", cmd)

	DatabaseClient.SetDatabase(database)

	fmt.Println(DatabaseClient.database)
	databaseClient := DatabaseClient.getDatabaseInstance()
	db := databaseClient.Database(DatabaseClient.database)
	runCommandResponse := db.RunCommand(context.TODO(), cmd)

	return runCommandResponse

}

func (DatabaseClient *DatabaseClient) ExecuteCommand(database string, cmd bson.D) map[string]interface{} {

	var cmdResponse map[string]interface{}
	//var bsonCommand bson.D

	//bson.Unmarshal(cmd, &bsonCommand)

	DatabaseClient.SetDatabase(database)

	fmt.Println(DatabaseClient.database)
	databaseClient := DatabaseClient.getDatabaseInstance()
	db := databaseClient.Database(DatabaseClient.database)
	runCommandErr := db.RunCommand(context.TODO(), cmd).Decode(&cmdResponse)

	if runCommandErr != nil {
		logger.LogError.Panicf("Exception while running execute command, err - %v", runCommandErr)
	}

	return cmdResponse
}

func (DatabaseClient *DatabaseClient) SetDatabase(database string) {
	DatabaseClient.database = database
}

func (DatabaseClient *DatabaseClient) GetDatabase() string {
	return DatabaseClient.database
}

func (DatabaseClient *DatabaseClient) GetCollection() string {
	return DatabaseClient.collection
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

	if err != nil {
		fmt.Errorf("Error while connecting to the federated client = ", err)
		panic(err)
	}
	response, _ := json.Marshal(client)

	fmt.Println(string(response))

	DatabaseClient.client = client

	if err != nil {
		panic(err)
	} else {
		return client
	}
}
