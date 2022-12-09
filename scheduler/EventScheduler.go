package scheduler

import (
	"mongomonitor/jobs/executor"
)

type JobScheduler struct {
	Time    string
	Payload map[string]string
}

func (js *JobScheduler) Schedule(Job executor.IJobExecutor) {
	Job.Execute(js.Payload)
	//c := gocron.NewScheduler()
	//c.Every(1).Days().At(js.Time).Do(Job)
	//c.Every(1).Seconds().Do(Job.Execute)
	//c.Start()
	//select {}
}
