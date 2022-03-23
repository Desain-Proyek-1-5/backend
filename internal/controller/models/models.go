package models

import "distancing-detect-backend/internal/entity"

type MqttAlert struct {
	Classroom       string `json:"class"`
	TotalViolations int    `json:"number_of_violations"`
	ImageLink       string `json:"photo_link"`
}

type ViolationData struct {
	Class           string `json:"class"`
	Time            string `json:"timestamp"`
	TotalViolations int    `json:"violation"`
	ImageLink       string `json:"photo_link"`
}

type TelegramPhoto struct {
	ChatID    int64  `json:"chat_id"`
	PhotoLink string `json:"photo"`
}

type TelegramOutgoingMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type TelegramIncomingMessage struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type TelegramWebhookSet struct {
	Url string `json:"url"`
}

func BuildViolationData(violation *entity.ViolationData) ViolationData {
	return ViolationData{
		Class:           violation.Class,
		Time:            violation.Time.Format("2006-01-02 15:04:05"),
		TotalViolations: violation.TotalViolations,
		ImageLink:       violation.ImageLink,
	}

}
