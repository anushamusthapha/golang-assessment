package main

import (
	"golang-assessment/config"
	"golang-assessment/logger"
	"golang-assessment/routers"
)

func main() {
	logger.InitLogger()
	db := config.SetupDatabase()
	router := routers.SetupRouter(db)
	logger.Log.Info("Starting the server on port 8080")
	router.Run(":8080")
}
