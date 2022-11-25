package scheduler

import "mongomonitor/jobs/executor"

type IJobScheduler interface {
	scheduler(Job *executor.IJobExecutor)
}
