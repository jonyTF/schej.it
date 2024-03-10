package gcloud

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"schej.it/server/logger"
)

func TestCreateEmailTask(t *testing.T) {
	// Init logfile
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Init logger
	logger.Init(logFile)

	// Load .env file
	err = godotenv.Load("../../.env")
	if err != nil {
		logger.StdErr.Panicln("Error loading .env file")
	}

	InitTasks()
	CreateEmailTask("schej.team@gmail.com", "Jonathan", "casablanca", "65e636bb760d3ea2e113e161")
}
