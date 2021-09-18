package main

import (
	"capstone/config"
	"capstone/pkg/database"
	"capstone/pkg/logger"
	"capstone/pkg/mqtt"
	"capstone/pkg/router"
	"log"
)

func main() {
	Logger := logger.NewLogger("log.txt")
	databaseConfig, err := config.LoadDatabaseConfiguration()
	if err != nil {
		log.Printf("Error setting database : %s\n", err.Error())
		return
	}
	router := router.NewRouterInstance()
	Database, err := database.NewDatabase("mysql",
		databaseConfig.Username, databaseConfig.Password, databaseConfig.Address,
		databaseConfig.DatabaseName)
	if err != nil {
		Logger.ErrorLogger.Println("Invalid database credentials supplied: ", err.Error())
	}
	mqtt := mqtt.NewMqttClient(Logger)
	router.Start()

}
