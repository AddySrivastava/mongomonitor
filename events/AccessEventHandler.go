package events

import (
	"mongomonitor/clients/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `bson:"principalName"`
	Source   string `bson:"remote"`
}

type AccessEvent struct {
	Message     string `bson:"msg"`
	AuthSuccess bool
	Client      User      `bson:"attr"`
	Severity    string    `bson:"s"`
	Timestamp   time.Time `bson:"t"`
}

type AccessEventHandler struct {
	EventDispatcher        *Dispatcher
	RepositoryClient       *repository.IRepositoryClient
	MongoMonitorRepository *repository.MongomonitorRepository
}

func (accessEventHandler *AccessEventHandler) Handle(payload []byte) (bool, error) {
	//set the variables inside payload
	var accessEvent AccessEvent

	repository := *accessEventHandler.RepositoryClient

	bson.UnmarshalExtJSON(payload, false, &accessEvent)

	if accessEvent.Message == "Authentication succeeded" {
		accessEvent.AuthSuccess = true
	} else {
		accessEvent.AuthSuccess = false
	}

	repository.SetDatabase("mongomonitor")
	repository.SetCollection("access")
	repository.InsertMany(bson.A{accessEvent})

	return true, nil

}
