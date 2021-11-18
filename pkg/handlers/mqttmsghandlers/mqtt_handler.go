package mqtthandler

import (
	"capstone/pkg/database"
	"capstone/pkg/handlers/telegram"
	"capstone/pkg/logger"
	"capstone/pkg/mqtt"
)

type MqttHandlers struct {
	Mqtt     *mqtt.MqttClient
	Database *database.DBInstance
	Logger   *logger.LoggerInstance
	Telegram *telegram.Telegram
}

func NewMqttHandler(MQTT *mqtt.MqttClient, Database *database.DBInstance, Logger *logger.LoggerInstance, Telegram *telegram.Telegram) *MqttHandlers {
	return &MqttHandlers{MQTT, Database, Logger, Telegram}
}

func (m *MqttHandlers) RegisterHandlers() {
	m.Mqtt.RegisterHandlerAndSubscribe("alert", 1, m.HandleAlert)
}
