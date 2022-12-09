package scheduler

import (
	"mongomonitor/events"
)

type IJobScheduler interface {
	scheduler(event *events.IEventHandler)
}
