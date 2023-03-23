package events

import (
	"fmt"
	"mongomonitor/clients/repository"
)

func CreateHandler(handlerType string, eventDispatcher *Dispatcher, databaseClient *repository.IRepositoryClient, mongoMonitor *repository.MongomonitorRepository) (IEventHandler, error) {
	switch handlerType {
	case "FIRE_ACCESS_EVENT":
		return &AccessEventHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient, MongoMonitorRepository: mongoMonitor}, nil
	case "INITIATE_LOG_FETCH":
		return &LogFetchHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient, MongoMonitorRepository: mongoMonitor}, nil
	case "INITIATE_LOG_UPLOAD":
		return &LogUploadHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient, MongoMonitorRepository: mongoMonitor}, nil
	case "INITIATE_SERVER_STATS_COLLECTOR":
		return &ServerStatsCollectorHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient, MongoMonitorRepository: mongoMonitor}, nil
	case "INITIATE_DB_STATS_COLLECTOR":
		return &DatabaseStatsHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient, MongoMonitorRepository: mongoMonitor}, nil
	case "INTIATE_CUSTOM_EVENT_COLLECTOR":
		return &CustomEventHandler{EventDispatcher: eventDispatcher, RepositoryClient: databaseClient, MongoMonitorRepository: mongoMonitor}, nil
	default:
		return nil, fmt.Errorf("Invalid handler type")
	}
}
