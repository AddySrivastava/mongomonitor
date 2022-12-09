package cmd

import (
	"log"
	"mongomonitor/clients/repository"
	"mongomonitor/events"
	"mongomonitor/factories"

	"github.com/spf13/cobra"
)

var url string
var time string

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Periodically monitor and store the MongoDB deployment",
	Run: func(cmd *cobra.Command, args []string) {

		//Initialize database connection

		const connectionString = "mongodb+srv://admin:passwordone@adityas-m10.4xwip.mongodb.net/?retryWrites=true&w=majority"
		options := make(map[string]string)

		databaseFactory := factories.DatabaseFactory{}

		var databaseClient repository.IRepositoryClient

		databaseClient = databaseFactory.CreateDatabase("database", connectionString, options)

		//Initialize the event dispatcher and Register Event Handlers
		eventDispatcher := &events.Dispatcher{Events: make(map[string]*events.IEventHandler)}

		accessEventHandler, _ := events.CreateHandler("FIRE_ACCESS_EVENT", eventDispatcher, &databaseClient)
		logFetchHandler, _ := events.CreateHandler("INITIATE_LOG_FETCH", eventDispatcher, &databaseClient)
		logUploadHandler, _ := events.CreateHandler("INITIATE_LOG_UPLOAD", eventDispatcher, &databaseClient)
		serverStatsCollector, _ := events.CreateHandler("INITIATE_SERVER_STATS_COLLECTOR", eventDispatcher, &databaseClient)

		eventDispatcher.Register(&accessEventHandler, "FIRE_ACCESS_EVENT")
		eventDispatcher.Register(&logFetchHandler, "INITIATE_LOG_FETCH")
		eventDispatcher.Register(&logUploadHandler, "INITIATE_LOG_UPLOAD")
		eventDispatcher.Register(&serverStatsCollector, "INITIATE_SERVER_STATS_COLLECTOR")

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
		logEventPayload := []byte(`{"path": "mongodbLog.log.gz"}`)

		//err := eventDispatcher.Dispatch("INITIATE_LOG_UPLOAD", &logEventPayload)
		err2 := eventDispatcher.Dispatch("INITIATE_SERVER_STATS_COLLECTOR", &logEventPayload)

		log.Fatal(err2)

		// call schedule with the log ingestion executor
		//scheduler.Schedule(jobExecutor)
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().StringVarP(&url, "url", "u", "", "Connection string eg - mongodb+srv://username:password@host")
	monitorCmd.Flags().StringVarP(&time, "time", "t", "00:00", "timestamp when the job will get executed in UTC 24h format eg- --time 18:00")
	monitorCmd.MarkFlagRequired("url")
}
