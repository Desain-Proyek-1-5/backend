package main

import (
	"distancing-detect-backend/pkg/database"
	httphandler "distancing-detect-backend/pkg/handlers/httpmsghandlers"
	mqtthandler "distancing-detect-backend/pkg/handlers/mqttmsghandlers"
	"distancing-detect-backend/pkg/handlers/telegram"
	"distancing-detect-backend/pkg/logger"
	"distancing-detect-backend/pkg/mqtt"
	"distancing-detect-backend/pkg/router"
	"fmt"
	"os"
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
	database, err := database.NewDatabase(os.Getenv("DATABASE_URL"))

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
