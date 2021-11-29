package models

type MqttAlert struct {
	Classroom       string `json:"class"`
	TotalViolations int    `json:"number_of_violations"`
	ImageLink       string `json:"photo_link"`
}

type ViolationData struct {
	Time            string `json:"time"`
	TotalViolations int    `json:"number_of_violations"`
	Class           string `json:"class"`
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
