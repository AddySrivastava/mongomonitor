package events

import (
	"fmt"
	"mongomonitor/clients/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Metrics struct {
	Count           int64 `bson:"count"`
	LatencyInMicros int64 `bson:"latencyInMicros"`
}

type SeverMetrics struct {
	Namespace   string
	Day         int
	Month       time.Month
	Year        int
	Hour        int
	Total       Metrics `bson:"total"`
	Queries     Metrics `bson:"queries"`
	Getmore     Metrics `bson:"getmore"`
	Insert      Metrics `bson:"insert"`
	Update      Metrics `bson:"update"`
	Commands    Metrics `bson:"commands"`
	CreatedDate time.Time
}

type ServerStatsCollectorHandler struct {
	EventDispatcher        *Dispatcher
	RepositoryClient       *repository.IRepositoryClient
	MongoMonitorRepository *repository.MongomonitorRepository
}

func (sch *ServerStatsCollectorHandler) Handle(payload []byte) (bool, error) {

	repository := *sch.RepositoryClient

	repository.SetDatabase("admin")

	//bsonCommand := []byte({serverStatus:1})
	byteCommand := bson.D{{Key: "top", Value: 1}}

	responseMap := repository.ExecuteCommand("admin", byteCommand)

	totals := responseMap["totals"]

	var serverMetrics SeverMetrics

	fmt.Println(totals)

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
		if err := bson.Unmarshal(bsonStr, &serverMetrics); err != nil {
			fmt.Println(err)
		}
		//fmt.Println(k)
		serverMetrics.Namespace = k

		t := time.Now()

		serverMetrics.CreatedDate = time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+1, t.Minute(), t.Second(), t.Nanosecond(), time.UTC)

		y, m, d := serverMetrics.CreatedDate.Date()

		serverMetrics.Day = d
		serverMetrics.Month = m
		serverMetrics.Year = y
		serverMetrics.Hour = serverMetrics.CreatedDate.Hour()

		//spew.Dump(serverMetrics)
		sch.MongoMonitorRepository.AddServerStats(serverMetrics)
	}

	return true, nil
}
