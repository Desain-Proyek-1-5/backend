package mqtt

import (
	"capstone/pkg/database"
	"capstone/pkg/logger"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	MqttInstance mqtt.Client
	Logger       *logger.LoggerInstance
}

type RequestHandler interface {
	Handle(client mqtt.Client, msg mqtt.Message)
}

func NewMqttClient(Logger *logger.LoggerInstance, Database *database.DBInstance) *MqttClient {
	var client mqtt.Client
	return &MqttClient{MqttInstance: client, Logger: Logger}
}

func (m *MqttClient) SetupMqttClient(Broker string, Port int, ClientID string) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", Broker, Port))
	opts.SetClientID(ClientID)
	opts.SetDefaultPublishHandler(m.messageHandler)
	opts.OnConnect = m.connectHandler
	opts.OnConnectionLost = m.connectLostHandler
	m.MqttInstance = mqtt.NewClient(opts)
	if token := m.MqttInstance.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (m *MqttClient) RegisterHandlerAndSubscribe(Topic string, Qos byte, Function RequestHandler) {
	m.MqttInstance.Subscribe(Topic, Qos, Function.Handle)
}

func (m *MqttClient) connectHandler(client mqtt.Client) {
	m.Logger.InfoLogger.Println("Connected into MQTT client")

}
func (m *MqttClient) connectLostHandler(client mqtt.Client, err error) {
	m.Logger.WarningLogger.Println("Disconnected from MQTT Broker: ", err.Error())

}

func (m *MqttClient) messageHandler(client mqtt.Client, msg mqtt.Message) {
	m.Logger.WarningLogger.Println("Unhandled message received: ", string(msg.Payload()))
}
