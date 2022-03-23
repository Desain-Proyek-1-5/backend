package controller

import (
	"bytes"
	"distancing-detect-backend/internal/controller/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Controller) sendTelegramMessage(message string, chatid int64) {
	msg := models.TelegramOutgoingMessage{
		ChatID: chatid,
		Text:   message,
	}
	req, err := json.Marshal(msg)
	if err != nil {
		c.handleInternalServerError(nil, err)
	}
	resp, err := http.Post(c.botUrl+"/sendMessage", "application/json", bytes.NewBuffer(req))
	if resp.StatusCode != 200 {
		c.handleInternalServerError(nil, errors.New("Telegram Not Acknowledging message"))
		return
	}
	if err != nil {
		c.handleInternalServerError(nil, err)
		return
	}
	c.logger.InfoLogger.Println("Sending message to " + fmt.Sprint(chatid))

}

func (c *Controller) sendTelegramImage(imageLink string, chatid int64) {
	msg := models.TelegramPhoto{
		ChatID:    chatid,
		PhotoLink: imageLink,
	}
	req, err := json.Marshal(msg)
	if err != nil {
		c.handleInternalServerError(nil, err)
	}
	resp, err := http.Post(c.botUrl+"/sendPhoto", "application/json", bytes.NewBuffer(req))
	if resp.StatusCode != 200 {
		c.handleInternalServerError(nil, errors.New("Telegram Not Acknowledging message"))
		return
	}
	if err != nil {
		c.handleInternalServerError(nil, err)
		return
	}

}

func (c *Controller) handleTelegramMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.TelegramIncomingMessage
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.handleInternalServerError(w, err)
		return
	}
	err = json.Unmarshal(req, &msg)
	if err != nil {
		c.handleInternalServerError(w, err)
		return
	}
	c.sendTelegramMessage("We have received your message: "+msg.Message.Text, msg.Message.Chat.ID)

}
