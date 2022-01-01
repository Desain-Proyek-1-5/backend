package mqtthandler

import (
	"capstone/pkg/models"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (m *MqttHandlers) HandleAlert(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("RECEIVED ALERT")
	timeStamp := time.Now()
	time := string(timeStamp.Format("2006-01-02 15:04:05"))
	var receivedAlert models.MqttAlert
	json.Unmarshal(msg.Payload(), &receivedAlert)
	query := fmt.Sprintf("INSERT INTO violations (Class, TotalViolations, Timeofdetection, photolink) VALUES ('%s',%d,'%s','%s');",
		receivedAlert.Classroom, receivedAlert.TotalViolations, time, receivedAlert.ImageLink)
	message := fmt.Sprintf("Terdeteksi Pelanggaran jaga jarak di kelas %s", receivedAlert.Classroom)
	fmt.Println(query)
	var wg sync.WaitGroup
	wg.Add(2)
	go m.saveToDatabase(&wg, query)
	go m.Telegram.SendTelegramMessage(&wg, message, receivedAlert.ImageLink, receivedAlert.Classroom)
	wg.Wait()
}

func (m *MqttHandlers) saveToDatabase(wg *sync.WaitGroup, query string) {
	err := m.Database.AddData(query)
	if err != nil {
		m.Logger.InfoLogger.Println("Error has occured , ", err.Error())
	}
	wg.Done()

}
