package main

import (
	"capstone/pkg/database"
	httphandler "capstone/pkg/handlers/httpmsghandlers"
	mqtthandler "capstone/pkg/handlers/mqttmsghandlers"
	"capstone/pkg/handlers/telegram"
	"capstone/pkg/logger"
	"capstone/pkg/mqtt"
	"capstone/pkg/router"
	"fmt"
)

func main() {
	logger := logger.NewLogger()
	/*
		databaseConfig, err := config.LoadDatabaseConfiguration()

		if err != nil {
			log.Printf("Error setting database : %s\n", err.Error())
			return
		}
	*/
	router := router.NewRouterInstance()
	database, err := database.NewDatabase("mysql",
		"root", "123jonathan123100300!!!", "localhost:3306",
		"testers")

	if err != nil {
		logger.ErrorLogger.Println("Invalid database credentials supplied: ", err.Error())
	}
	telegram := telegram.NewTelegram(logger)
	mqtt := mqtt.NewMqttClient(logger)
	mqtt.SetupMqttClient("broker.emqx.io", 1883, "golang1231")
	fmt.Println("Connected to MQTT Client")
	httpHandler := httphandler.NewHTTPHandler(router, database, logger)
	httpHandler.RegisterHandlers()
	mqttHandler := mqtthandler.NewMqttHandler(mqtt, database, logger, telegram)
	mqttHandler.RegisterHandlers()
	router.Start()

}
