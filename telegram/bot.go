package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	telegramBotAPI = "https://api.telegram.org/bot%s/%s"
	sendMessage    = "sendMessage"
)

type Message struct {
	BotToken string `json:"-"`
	ChatID   string `json:"chat_id"`
	Text     string `json:"text"`
}

func (m Message) url() string {
	return fmt.Sprintf(telegramBotAPI, m.BotToken, sendMessage)
}

func (m Message) Send() error {
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	_, err = http.Post(m.url(), "application/json", bytes.NewBuffer(body))
	return err
}
