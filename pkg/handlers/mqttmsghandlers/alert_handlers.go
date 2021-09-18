package mqtthandlers

import (
	"capstone/pkg/database"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type AlertHandler struct {
	Mqtt     mqtt.Client
	Database *database.DBInstance
}
type Alert struct {
	Classroom        string `json:"class"`
	ExpectedDistance string `json:"distance"`
}

func (a *AlertHandler) Handle(client mqtt.Client, msg mqtt.Message) {
	timeStamp := time.Now().String()
	var receivedAlert Alert
	json.Unmarshal(msg.Payload(), &receivedAlert)
	a.Database.AddData(fmt.Sprintf("INSERT INTO %s (Distance, Time) VALUES (%s,%s)",
		receivedAlert.Classroom, receivedAlert.ExpectedDistance, timeStamp))
}
