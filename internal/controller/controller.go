package controller

import (
	"bytes"
	"distancing-detect-backend/internal/controller/models"
	"distancing-detect-backend/internal/usecase"
	"distancing-detect-backend/pkg/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Controller struct {
	httpRouter       *mux.Router
	mqttRouter       mqtt.Client
	logger           *logger.LoggerInstance
	service          *usecase.Usecase
	telegramChannels map[string]int64
	mqttChannel      string
	botUrl           string
}

func NewController(mqttBroker string, mqttClientId string, mqttChannel string, telegramChannels map[string]int64,
	botUrl string, usecase *usecase.Usecase, logger *logger.LoggerInstance) *Controller {
	router := mux.NewRouter().StrictSlash(true)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID(mqttClientId)
	controller := &Controller{
		httpRouter:       router,
		mqttRouter:       mqtt.NewClient(opts),
		service:          usecase,
		telegramChannels: telegramChannels,
		botUrl:           botUrl,
		logger:           logger,
	}
	controller.registerHttpHandler()
	controller.registerMqttHandler()
	controller.setTelegramWebHook()
	return controller
}

func (c *Controller) registerHttpHandler() {
	c.httpRouter.HandleFunc("/{class}", c.GetViolationsOfClass).Methods("POST")
	c.httpRouter.HandleFunc("/", c.GetAllViolations).Methods("GET")
	c.httpRouter.HandleFunc("/telegram/incoming", c.handleTelegramMessage).Methods("POST")
}

func (c *Controller) registerMqttHandler() {
	c.mqttRouter.Subscribe("3deef803-2854-495d-b641-677c7bda1979", 1, c.HandleAlert)
}

func (c *Controller) setTelegramWebHook() {
	callbackUrl := os.Getenv("HEROKU_URL") + "telegram/incoming"
	msg := models.TelegramWebhookSet{
		Url: callbackUrl,
	}
	req, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.botUrl)
	resp, err := http.Post(c.botUrl+"/setWebhook", "application/json", bytes.NewBuffer(req))
	if resp.StatusCode != 200 {
		c.logger.ErrorLogger.Println("Telegram Error: ", resp.Body)
		panic("Telegram Error!")
	}
	if err != nil {
		panic(err)
	}

}

func (c *Controller) Start() {
	cors := cors.AllowAll()
	port := os.Getenv("PORT")
	handler := cors.Handler(c.httpRouter)
	if token := c.mqttRouter.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	http.ListenAndServe(":"+port, handler)
}
