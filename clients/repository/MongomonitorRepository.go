package repository

import (
	fmt "fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongomonitorRepository struct {
	RepositoryClient *IRepositoryClient
}

func (mongomonitorRepository *MongomonitorRepository) UpsertDBStats(namespace string, createdDate time.Time, day int, month int, year int, docs interface{}) {

	repository := *mongomonitorRepository.RepositoryClient

	repository.SetDatabase("mongomonitor")
	repository.SetCollection("dbStats")

	dayPartUpdate := fmt.Sprintf("stats.%d", day)

	filter := bson.D{{"namespace", namespace}, {"month", month}, {"year", year}, {"day", day}}
	update := bson.D{{"$set", bson.D{{dayPartUpdate, docs}}}, {"$setOnInsert", bson.D{{"created_date", createdDate}}}}
	options := options.Update().SetUpsert(true)

	bsonOptions, _ := bson.Marshal(options)
	repository.Upsert(filter, update, bsonOptions)
}

func (mongomonitorRepository *MongomonitorRepository) AddDBStats(docs interface{}) {

	repository := *mongomonitorRepository.RepositoryClient

	repository.SetDatabase("mongomonitor")
	repository.SetCollection("dbStats")

	repository.InsertOne(docs)
}

func (mongomonitorRepository *MongomonitorRepository) AddServerStats(docs interface{}) {
	repository := *mongomonitorRepository.RepositoryClient

	repository.SetDatabase("mongomonitor")
	repository.SetCollection("metrics")

	repository.InsertOne(docs)
}
