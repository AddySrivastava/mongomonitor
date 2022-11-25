package cmd

import (
	"log"
	"mongomonitor/jobs/executor"
	"mongomonitor/jobs/scheduler"

	"github.com/spf13/cobra"
)

var url string
var time string

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Periodically monitor and store the MongoDB deployment",
	Run: func(cmd *cobra.Command, args []string) {

		//url, urlErr := cmd.LocalFlags().GetString("url")
		time, timeErr := cmd.LocalFlags().GetString("time")

		//if urlErr != nil {
		//	log.Fatal("Not a valid URL", urlErr)
		//}

		if timeErr != nil {
			log.Fatal("Not a valid time value", timeErr)
		}

		// create the scheduler object to schedule the jobs based on the time specified

		scheduler := scheduler.JobScheduler{
			Time: time,
		}

		// create the job executor that ingests the logs based on the function

		LogIngestionExecutor := &executor.LogIngestionExecutor{}

		// call schedule with the log ingestion executor
		scheduler.Schedule(LogIngestionExecutor)
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().StringVarP(&url, "url", "u", "", "Connection string eg - mongodb+srv://username:password@host")
	monitorCmd.Flags().StringVarP(&time, "time", "t", "00:00", "timestamp when the job will get executed in UTC 24h format eg- --time 18:00")
	monitorCmd.MarkFlagRequired("url")
}
