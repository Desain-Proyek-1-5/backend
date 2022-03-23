package main

import (
	"distancing-detect-backend/internal/controller"
	"distancing-detect-backend/internal/entity"
	"distancing-detect-backend/internal/repository"
	"distancing-detect-backend/internal/usecase"
	"distancing-detect-backend/pkg/logger"
	"encoding/json"
	"io/ioutil"
	"os"
)

func parseConfig() (string, map[string]int64, string, string) {
	/*
		Returns Telegram Bot id, class channel as a map, mqtt notification channel id, and mqtt id
	*/
	var config entity.Configuration
	configFile, err := ioutil.ReadFile("../configs/config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
	var classChannels map[string]int64
	for _, class := range config.Telegram.Class {
		classChannels[class.Classroom] = class.GroupId
	}
	return config.Telegram.BotId, classChannels, config.Mqtt.MqttNotificationChannel, config.Mqtt.MqttId
}
func main() {
	logger := logger.NewLogger()
	telegramBotUrl, classChannels, mqttChannel, mqttId := parseConfig()
	repo := repository.NewRepository(os.Getenv("DATABASE_URL"))
	service := usecase.NewService(*repo)
	controller := controller.NewController("tcp://broker.emqx.io:1883", mqttId, mqttChannel, classChannels, telegramBotUrl, service, logger)
	controller.Start()
}
