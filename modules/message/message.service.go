package message

import (
	"encoding/json"
	"io"
)

type MessageService struct {
	MessageRepository
}
type CreateMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (s *MessageService) Create(body io.Reader) {
	decoder := json.NewDecoder(body)
	var message CreateMessage
	decoder.Decode(&message)
	s.MessageRepository.Create(message.Title, message.Text)
}
