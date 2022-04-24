package message

import (
	"encoding/json"
	"io"
)

type MessageService struct {
	MessageRepository
}

func (s *MessageService) Create(body io.Reader) Message {
	decoder := json.NewDecoder(body)
	var message Message
	decoder.Decode(&message)
	s.MessageRepository.Create(message.Title, message.Text)
	return message
}
func (s *MessageService) GetAll() []Message {

	return s.MessageRepository.GetAll()
}
