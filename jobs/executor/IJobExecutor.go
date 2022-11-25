package executor

import "mongomonitor/types"

type IJobExecutor interface {
	Execute() types.JobExecutionResponse
}
