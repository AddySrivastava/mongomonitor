package main

import (
	"log"
	cmd "mongomonitor/cmd"
	"mongomonitor/utils"
	"os"

	godotenv "github.com/joho/godotenv"
)

var logger = utils.GetLogger()

func main() {
	envFileLoadErr := godotenv.Load()

	if envFileLoadErr != nil {
		log.Fatal("Something went wrong while loading env var file")
	}

	// Read config file and pass it along the executor as a payload for the handlers to process
	data, readFileErr := os.ReadFile("./mongomonitorConfig.yaml")

	if readFileErr != nil {
		logger.LogError.Panicf("Could not open configuration file, Error: %#v", readFileErr)
	}

	cmd.Execute(data)
}
