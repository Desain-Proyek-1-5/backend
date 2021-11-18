package mqtt

import (
	"capstone/pkg/logger"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	MqttInstance mqtt.Client
	Logger       *logger.LoggerInstance
}

func NewMqttClient(Logger *logger.LoggerInstance) *MqttClient {
	var client mqtt.Client
	return &MqttClient{MqttInstance: client, Logger: Logger}
}

func (m *MqttClient) SetupMqttClient(Broker string, Port int, ClientID string) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://broker.emqx.io:1883")
	opts.SetClientID(ClientID)
	opts.SetDefaultPublishHandler(m.messageHandler)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	opts.OnConnect = m.connectHandler
	opts.OnConnectionLost = m.connectLostHandler
	m.MqttInstance = mqtt.NewClient(opts)
	if token := m.MqttInstance.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	m.Logger.InfoLogger.Printf("Connected to: %s\n", Broker)
}

func (m *MqttClient) RegisterHandlerAndSubscribe(Topic string, Qos byte, Function func(client mqtt.Client, msg mqtt.Message)) {
	token := m.MqttInstance.Subscribe(Topic, Qos, Function)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
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
