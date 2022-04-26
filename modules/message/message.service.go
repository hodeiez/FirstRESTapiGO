package message

import (
	"encoding/json"
	"io"
)

type MessageService struct {
	MessageRepository
}

var message Message

func (s *MessageService) Create(body io.Reader) Message {
	decoder := json.NewDecoder(body)
	decoder.Decode(&message)
	s.MessageRepository.Create(message.Title, message.Text)
	return message
}
func (s *MessageService) GetAll() []Message {

	return s.MessageRepository.GetAll()
}
