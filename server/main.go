package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"schej.it/server/db"
	"schej.it/server/logger"
	"schej.it/server/routes"
	"schej.it/server/slackbot"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "schej.it/server/docs"
)

// @title Schej.it API
// @version 1.0
// @description This is the API for Schej.it!

// @host localhost:3002

func main() {
	// Set release flag
	release := flag.Bool("release", false, "Whether this is the release version of the server")
	flag.Parse()
	if *release {
		os.Setenv("GIN_MODE", "release")
		gin.SetMode(gin.ReleaseMode)
	} else {
		os.Setenv("GIN_MODE", "debug")
	}

	// Init logfile
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// Init logger
	logger.Init(logFile)

	// Load .env variables
	loadDotEnv()

	// Init router
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}
		return fmt.Sprintf("%v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "https://www.schej.it", "https://schej.it"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init database
	closeConnection := db.Init()
	defer closeConnection()

	// Session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	// Init routes
	routes.InitAuth(router)
	routes.InitUser(router)
	routes.InitEvents(router)
	routes.InitUsers(router)
	routes.InitAnalytics(router)
	slackbot.InitSlackbot(router)

	// Init swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run server
	router.Run(":3002")
}

// Load .env variables
func loadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		logger.StdErr.Panicln("Error loading .env file")
	}
}
