package events

import (
	"fmt"
	"time"

	"mongomonitor/clients/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Metrics struct {
	Count           int64 `bson:"count"`
	LatencyInMicros int64 `bson:"latencyInMicros"`
}

type OpCounterStats struct {
	Total    Metrics `bson:"total"`
	Queries  Metrics `bson:"queries"`
	Getmore  Metrics `bson:"getmore"`
	Insert   Metrics `bson:"insert"`
	Update   Metrics `bson:"update"`
	Commands Metrics `bson:"commands"`
}

type SeverMetrics struct {
	Namespace   string
	Day         int
	Month       time.Month
	Year        int
	Stats       map[int]OpCounterStats
	CreatedDate time.Time
}

type ServerStatsCollectorHandler struct {
	EventDispatcher  *Dispatcher
	RepositoryClient *repository.IRepositoryClient
}

func (sch *ServerStatsCollectorHandler) Handle(payload []byte) (bool, error) {

	repository := *sch.RepositoryClient

	repository.SetDatabase("admin")

	//bsonCommand := []byte({serverStatus:1})
	cmd := bson.D{{"top", 1}}
	byteCommand, _ := bson.Marshal(cmd)

	responseMap := repository.ExecuteCommand(byteCommand)

	totals := responseMap["totals"]

	var opcounterStat OpCounterStats
	var serverMetrics SeverMetrics

	for k, v := range totals.(map[string]interface{}) {
		//var collStats OpCounterStats
		if k == "note" {
			continue
		}

		bsonStr, err := bson.Marshal(v)
		if err != nil {
			fmt.Println(err)
		}

		// Convert json string to struct
		if err := bson.Unmarshal(bsonStr, &opcounterStat); err != nil {
			fmt.Println(err)
		}

		serverMetrics.Namespace = k

		serverMetrics.Stats = make(map[int]OpCounterStats)

		t := time.Now()

		serverMetrics.CreatedDate = time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+1, t.Minute(), t.Second(), t.Nanosecond(), time.UTC)

		y, m, d := serverMetrics.CreatedDate.Date()
		hourPart := serverMetrics.CreatedDate.Hour()

		serverMetrics.Stats[hourPart+1] = opcounterStat

		serverMetrics.Day = d
		serverMetrics.Month = m
		serverMetrics.Year = y

		repository.SetDatabase("mongomonitor")
		repository.SetCollection("metrics")

		hourPartUpdate := fmt.Sprintf("stats.%d", hourPart)

		filter := bson.D{{"namespace", serverMetrics.Namespace},
			{"day", serverMetrics.Day},
			{"month", serverMetrics.Month},
			{"year", serverMetrics.Year}}
		update := bson.D{{"$set", bson.D{{hourPartUpdate, opcounterStat}}}, {"$setOnInsert", bson.D{{"created_date", serverMetrics.CreatedDate}}}}
		opts := options.Update().SetUpsert(true)

		optsBytes, _ := bson.Marshal(opts)

		repository.Upsert(filter, update, optsBytes)
	}

	// if rec, ok := totals.([]OpCounterStats); ok {
	// 	for _, v := range rec {
	// 		fmt.Println(v)
	// 	}
	// }

	return true, nil
}
