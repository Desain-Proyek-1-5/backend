package entity

import "time"

type ViolationData struct {
	Class           string
	Time            time.Time
	TotalViolations int
	ImageLink       string
}

type MqttConfigurations struct {
	MqttId                  string `json:"mqtt_id"`
	MqttNotificationChannel string `json:"notification_channel"`
}

type TelegramConfigurations struct {
	BotId string         `json:"telegram_bot_id"`
	Class []ClassChannel `json:"classes"`
}

type ClassChannel struct {
	Classroom string `json:"class"`
	GroupId   int64  `json:"group_id"`
}

type Configuration struct {
	Telegram TelegramConfigurations `json:"telegram"`
	Mqtt     MqttConfigurations     `json:"mqtt"`
}

func NewViolation(class string, time time.Time, totalViolations int, imageLink string) *ViolationData {
	return &ViolationData{
		Class:           class,
		Time:            time,
		TotalViolations: totalViolations,
		ImageLink:       imageLink,
	}
}
