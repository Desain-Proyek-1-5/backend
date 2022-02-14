package mqtthandler

import (
	"distancing-detect-backend/pkg/database"
	"distancing-detect-backend/pkg/handlers/telegram"
	"distancing-detect-backend/pkg/logger"
	"distancing-detect-backend/pkg/mqtt"
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
	m.Mqtt.RegisterHandlerAndSubscribe("message_to_server_192213", 1, m.HandleAlert)
}
