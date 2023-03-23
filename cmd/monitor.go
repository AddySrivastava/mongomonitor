package cmd

import (
	"fmt"
	"log"
	"mongomonitor/clients/repository"
	"mongomonitor/events"
	"mongomonitor/factories"
	"os"
)

var url string
var time string

func Execute(payload []byte) {

	var databaseClient repository.IRepositoryClient

	//read connection string from environment varibales
	var connectionString = os.Getenv("MONGODB_URI")

	fmt.Println(connectionString)

	options := make(map[string]string)

	databaseFactory := (factories.DatabaseFactory{})

	databaseClient = databaseFactory.CreateDatabase("database", connectionString, options)

	mongoMonitorRepo := repository.MongomonitorRepository{RepositoryClient: &databaseClient}

	//Initialize the event dispatcher and Register Event Handlers
	eventDispatcher := &events.Dispatcher{Events: make(map[string]*events.IEventHandler)}

	accessEventHandler, _ := events.CreateHandler("FIRE_ACCESS_EVENT", eventDispatcher, &databaseClient, &mongoMonitorRepo)
	logFetchHandler, _ := events.CreateHandler("INITIATE_LOG_FETCH", eventDispatcher, &databaseClient, &mongoMonitorRepo)
	logUploadHandler, _ := events.CreateHandler("INITIATE_LOG_UPLOAD", eventDispatcher, &databaseClient, &mongoMonitorRepo)
	serverStatsCollector, _ := events.CreateHandler("INITIATE_SERVER_STATS_COLLECTOR", eventDispatcher, &databaseClient, &mongoMonitorRepo)
	databaseStatsCollector, _ := events.CreateHandler("INITIATE_DB_STATS_COLLECTOR", eventDispatcher, &databaseClient, &mongoMonitorRepo)
	customEventStatsCollector, _ := events.CreateHandler("INTIATE_CUSTOM_EVENT_COLLECTOR", eventDispatcher, &databaseClient, &mongoMonitorRepo)

	eventDispatcher.Register(&accessEventHandler, "FIRE_ACCESS_EVENT")
	eventDispatcher.Register(&logFetchHandler, "INITIATE_LOG_FETCH")
	eventDispatcher.Register(&logUploadHandler, "INITIATE_LOG_UPLOAD")
	eventDispatcher.Register(&serverStatsCollector, "INITIATE_SERVER_STATS_COLLECTOR")
	eventDispatcher.Register(&databaseStatsCollector, "INITIATE_DB_STATS_COLLECTOR")
	eventDispatcher.Register(&customEventStatsCollector, "INTIATE_CUSTOM_EVENT_COLLECTOR")

	//time, timeErr := cmd.LocalFlags().GetString("time")

	//if timeErr != nil {
	//	log.Fatal("Not a valid time value", timeErr)
	//}

	// create the scheduler object to schedule the jobs based on the time specified
	/*scheduler := scheduler.JobScheduler{
		Time:    time,
		Payload: make(map[string]string),
	}*/

	// create the job executor that ingests the logs based on the function
	//logEventPayload := []byte(`{"path": "mongodbLog.log.gz"}`)

	//err := eventDispatcher.Dispatch("INITIATE_LOG_UPLOAD", &logEventPayload)
	//err2 := eventDispatcher.Dispatch("INITIATE_SERVER_STATS_COLLECTOR", &logEventPayload)

	//("INITIATE_DB_STATS_COLLECTOR", &logEventPayload)

	err4 := eventDispatcher.Dispatch("INTIATE_CUSTOM_EVENT_COLLECTOR", &payload)

	//log.Fatal(err3)
	log.Fatal(err4)

	// call schedule with the log ingestion executor
	//scheduler.Schedule(jobExecutor)
}
