package events

import (
	"encoding/json"
	"fmt"
	"mongomonitor/clients/mongo"
	"mongomonitor/clients/repository"
	"os"
	"strings"
)

type LogFetchHandler struct {
	EventDispatcher        *Dispatcher
	RepositoryClient       *repository.IRepositoryClient
	MongoMonitorRepository *repository.MongomonitorRepository
}

func (logFetchHandler *LogFetchHandler) Handle(payload []byte) (bool, error) {

	var value map[string]interface{}

	// Open and create the file in read and write mode
	logFile, _ := os.OpenFile("mongodbLog.log.gz", os.O_CREATE|os.O_RDWR, 0755)

	//Get the environment variables for the call
	publicKey := os.Getenv("ATLAS_PUBLIC_KEY")
	privateKey := os.Getenv("ATLAS_PRIVATE_KEY")
	projectId := os.Getenv("ATLAS_PROJECT_ID")
	//clusterName := os.Getenv("ATLAS_CLUSTER_NAME")

	atlasClient := mongo.AtlasClient{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	hosts := atlasClient.GetPrimaryHostByProjects(projectId)

	json.Unmarshal(hosts, &value)

	fmt.Println(value)

	results := value["results"]

	if rec, ok := results.([]interface{}); ok {
		for _, v := range rec {
			if obj, ok := v.(map[string]interface{}); ok {
				_, after, found := strings.Cut(obj["userAlias"].(string), "adityas-m10")

				if (found == true) && (after[:6] == "-shard") && (obj["typeName"] == "REPLICA_PRIMARY") {
					logs := atlasClient.GetLogsWithRange(projectId, obj["hostname"].(string), "mongodb.gz", 1, 2)

					logFile.Write(logs)
				}
			}
		}

	} else {
		fmt.Println(rec)
	}

	return true, nil
}
