package utils

import (
	"log"
	"os"
)

type LoggerType struct {
	LogError *log.Logger
	LogWarn  *log.Logger
	LogInfo  *log.Logger
}

var LoggerInstance = LoggerType{}

func GetLogger() LoggerType {

	if LoggerInstance.LogError != nil && LoggerInstance.LogInfo != nil && LoggerInstance.LogWarn != nil {
		return LoggerInstance
	}

	file, fileOpenErr := os.OpenFile("mongmonitor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if fileOpenErr != nil {
		log.Fatal("Something went wrong while opening log file")
	}

	LoggerInstance.LogError = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerInstance.LogInfo = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerInstance.LogWarn = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	return LoggerInstance
}
