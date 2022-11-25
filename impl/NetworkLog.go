package impl

import (
	"time"
)

type operatingSystem struct {
	osType string
	architecture string
}

type NetworkLog struct {
	id int32
	message string
	os operatingSystem
	timestamp time.Time
	application string
	severity string
}

