package main

import (
	"distancing-detect-backend/internal/controller"
	"distancing-detect-backend/internal/entity"
	"distancing-detect-backend/internal/repository"
	"distancing-detect-backend/internal/usecase"
	"distancing-detect-backend/pkg/logger"
	"encoding/json"
	"os"
)

func parseConfig() (string, map[string]int64, string, string) {
	/*
		Returns Telegram Bot id, class channel as a map, mqtt notification channel id, and mqtt id
	*/
	configuration := `{
		"telegram": 
		{
			"telegram_bot_id" : "https://api.telegram.org/bot5120282372:AAFcXdDmwC4vsoERGFt797I0oQZAkSovMOA",
			"classes" : [
				{
					"class" : "IPA 1",
					"group_id" : -3498238
				},
				{
					"class" : "IPA 2",
					"group_id" : -2391239
				}
			]
		},
		"mqtt" :
		{
			"mqtt_broker" : "tcp://broker.emqx.io:1883",
			"mqtt_id" : "mqtt-http-server-uQkndf-1",
			"notification_channel" : "notification_channel_uQkndf"
		}
	}
	`
	var config entity.Configuration
	err := json.Unmarshal([]byte(configuration), &config)
	if err != nil {
		panic(err)
	}
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
