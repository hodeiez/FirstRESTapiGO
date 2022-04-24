package message

import (
	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func (repo *MessageRepository) Create(title string, message string) {
	repo.DB.Create(&Message{Title: title, Text: message})

}
func (repo *MessageRepository) GetAll() []Message {
	var messages []Message
	repo.DB.Find(&messages)
	return messages

}
