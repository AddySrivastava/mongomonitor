package events

import (
	"fmt"
	"mongomonitor/clients/repository"
)

func CreateHandler(handlerType string, eventDispatcher *Dispatcher, databaseClient *repository.IRepositoryClient) (IEventHandler, error) {
	switch handlerType {
	case "FIRE_ACCESS_EVENT":
		return &AccessEventHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient}, nil
	case "INITIATE_LOG_FETCH":
		return &LogFetchHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient}, nil
	case "INITIATE_LOG_UPLOAD":
		return &LogUploadHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient}, nil
	case "INITIATE_SERVER_STATS_COLLECTOR":
		return &ServerStatsCollectorHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient}, nil
	default:
		return nil, fmt.Errorf("Invalid handler type")
	}
}
