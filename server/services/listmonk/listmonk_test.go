package listmonk

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"schej.it/server/logger"
)

func TestSendEmail(t *testing.T) {
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

	SendEmail("schej.team@gmail.com", 8, bson.M{
		"eventName": "casablanca",
		"eventUrl":  "http://localhost:8080/e/65e636bb760d3ea2e113e161",
	})
}
