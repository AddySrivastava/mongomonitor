package events

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"mongomonitor/clients/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Database struct {
	Name string `bson:"name"`
}

type Collection struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
}

type CollectionStats struct {
	Namespace      string
	Day            int       `bson:"day"`
	Month          int       `bson:"month"`
	Year           int       `bson:"year"`
	Week           int       `bson:"week"`
	Count          int32     `bson:"count"`
	AverageObjSize int32     `bson:"avgObjSize"`
	Size           int32     `bson:"size"`
	TotalIndexSize int32     `bson:"totalIndexSize"`
	CreatedDate    time.Time `bson:"created_date"`
}

type DatabaseStatsHandler struct {
	EventDispatcher        *Dispatcher
	MongoMonitorRepository *repository.MongomonitorRepository
	RepositoryClient       *repository.IRepositoryClient
}

func (databaseStatsHandler *DatabaseStatsHandler) Handle(payload []byte) (bool, error) {
	var database []Database

	//get the databases

	repository := *databaseStatsHandler.RepositoryClient

	byteCommand := bson.D{{"listDatabases", 1}}

	responseMap := repository.ExecuteCommand("admin", byteCommand)

	databases := responseMap["databases"]

	bsonDatabase, _ := json.Marshal(&databases)

	json.Unmarshal(bsonDatabase, &database)

	//get the collections

	for _, vD := range database {
		if len(vD.Name) > 0 {
			var coll []Collection
			var collStats CollectionStats

			byteCommand := bson.D{{"listCollections", 1}}

			responseMap := repository.ExecuteCommand(vD.Name, byteCommand)

			cursor := responseMap["cursor"]
			firstBatch := (cursor.(map[string]interface{}))["firstBatch"]

			firstBatchByte, _ := json.Marshal(firstBatch)

			json.Unmarshal(firstBatchByte, &coll)

			for _, v := range coll {
				//run command

				if v.Type == "view" {
					continue
				}

				repository.SetDatabase(vD.Name)

				dbStatsCmd := bson.D{{"collStats", v.Name}}

				responseMap := repository.ExecuteCommand(vD.Name, dbStatsCmd)

				//spew.Dump(responseMap)
				if _, ok := responseMap["count"]; ok {
					collStats.Count = responseMap["count"].(int32) + rand.Int31n(100)

				}

				if _, ok := responseMap["totalIndexSize"]; ok {
					collStats.TotalIndexSize = responseMap["totalIndexSize"].(int32) + rand.Int31n(100)

				}

				if _, ok := responseMap["size"]; ok {
					collStats.Size = responseMap["size"].(int32) + rand.Int31n(100)

				}

				if _, ok := responseMap["avgObjSize"]; ok {
					collStats.AverageObjSize = responseMap["avgObjSize"].(int32) + rand.Int31n(100)

				}

				collStats.Namespace = fmt.Sprintf("%s.%s", vD.Name, v.Name)
				t := time.Now()
				collStats.CreatedDate = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
				collStats.Day = rand.Intn(30)   //+ t.Day()
				collStats.Month = rand.Intn(12) // + int(t.Month())
				collStats.Year = t.Year()
				collStats.Week = rand.Intn(5) // + (t.Weekday())
				databaseStatsHandler.MongoMonitorRepository.AddDBStats(collStats)

			}

		}
	}

	return true, nil
}
