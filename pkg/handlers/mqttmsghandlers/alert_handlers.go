package mqtthandler

import (
	"capstone/pkg/models"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (m *MqttHandlers) HandleAlert(client mqtt.Client, msg mqtt.Message) {
	timeStamp := time.Now().String()
	var receivedAlert models.MqttAlert
	json.Unmarshal(msg.Payload(), &receivedAlert)
	m.Database.AddData(fmt.Sprintf("INSERT INTO violations (Class, Distance, Time) VALUES ('%s','%s','%s')",
		receivedAlert.Classroom, receivedAlert.ExpectedDistance, timeStamp))
}
