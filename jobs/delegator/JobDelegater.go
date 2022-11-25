package delegator

import (
	"mongomonitor/jobs/executor"
	"mongomonitor/types"
)

func JobDelegator(jobType string, payload map[string]string) types.JobExecutionResponse {
	var JobExecutor executor.IJobExecutor

	switch jobType {
	case "ingest":
		JobExecutor = &executor.LogIngestionExecutor{}
	}

	response := JobExecutor.Execute()

	return response
}
