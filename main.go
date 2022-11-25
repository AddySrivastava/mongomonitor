package main

import (
	json "encoding/json"
	fmt "fmt"
	repository "mongomonitor/clients/repository"
	cmd "mongomonitor/cmd"
	factories "mongomonitor/factories"
	bson "go.mongodb.org/mongo-driver/bson"
)

func main() {
	const connectionString = "mongodb+srv://admin:passwordone@adityas-m10.4xwip.mongodb.net/?retryWrites=true&w=majority"
	options := make(map[string]string)

	databaseFactory := factories.DatabaseFactory{}

	var databaseClient repository.IRepositoryClient

	databaseClient = databaseFactory.CreateDatabase("database", connectionString, options)

	payload, _ := json.Marshal(databaseClient)

	fmt.Println(string(payload))
	databaseClient.SetDatabase("dummy")
	databaseClient.SetCollection("zips")

	docs := bson.A{bson.D{{"foo", "bar"}}}

	databaseClient.Insert(docs)

	cmd.Execute()

}
