package events

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"mongomonitor/clients/repository"
	"os"
)

type LogUpload struct {
	Path string `json:"path"`
}

type LogEventPayload struct {
	EventType string `json:"c"`
}

/*
Struct to upload log files
*/

type LogUploadHandler struct {
	EventDispatcher        *Dispatcher
	RepositoryClient       *repository.IRepositoryClient
	MongoMonitorRepository *repository.MongomonitorRepository
}

func (lu *LogUploadHandler) Handle(payload []byte) (bool, error) {

	var logEventPaylod LogEventPayload
	var logUpload LogUpload

	json.Unmarshal(payload, &logUpload)
	//Read the Log file

	logFile, err := os.OpenFile(logUpload.Path, os.O_RDONLY, 0755)
	defer logFile.Close()

	if err != nil {
		fmt.Println(err)

		return false, fmt.Errorf("Failed to open file - %s", err)
	}

	rawContents, err := gzip.NewReader(logFile)
	defer rawContents.Close()

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fileScanner := bufio.NewScanner(rawContents)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		json.Unmarshal(line, &logEventPaylod)
		if logEventPaylod.EventType == "ACCESS" {
			lu.EventDispatcher.Dispatch("FIRE_ACCESS_EVENT", &line)
		}
	}

	return true, nil
}
