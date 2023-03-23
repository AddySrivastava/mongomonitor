package events

import (
	"fmt"
	"mongomonitor/clients/repository"
	"mongomonitor/types"
	"mongomonitor/utils"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v2"
)

type CustomEventHandler struct {
	EventDispatcher        *Dispatcher
	MongoMonitorRepository *repository.MongomonitorRepository
	RepositoryClient       *repository.IRepositoryClient
}

type CustomEventPayload struct {
	Command string // command to run
	Metric  string // metric to monitor
}

type CustomEventConfigType struct {
	CustomMetrics struct {
		CollStats struct {
			CollectionName string   `yaml:"collectionName"`
			Metrics        []string `yaml:"metrics"`
		} `yaml:"collStats"`
		Top []string `yaml:"top"`
	} `yaml:"customMetrics"`
}

func (customEventHandler *CustomEventHandler) Handle(payload []byte) (bool, error) {

	// initialize logger
	var loggerInstance = utils.GetLogger()

	//repository := *customEventHandler.RepositoryClient

	//Load the config file

	var customEventConfig CustomEventConfigType

	fileUnmarshalError := yaml.Unmarshal(payload, &customEventConfig)

	if fileUnmarshalError != nil {
		loggerInstance.LogError.Panicf("Failed to parse the configuration file, Error: %#v", fileUnmarshalError)
	}

	fmt.Printf("%v", customEventConfig.CustomMetrics.CollStats.CollectionName)
	fmt.Printf("%v", customEventConfig.CustomMetrics.CollStats.Metrics)

	//Get Data for CollStats
	repository := *customEventHandler.RepositoryClient

	collStatsCommand := bson.D{{Key: "collStats", Value: "companies"}}

	var collStats types.TCollStats

	response := repository.ExecuteCommand2("sample_training", collStatsCommand)

	response.Decode(&collStats)

	return true, nil

}
