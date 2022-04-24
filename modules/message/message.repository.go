package message

import (
	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func (repo *MessageRepository) Create(title string, text string) Message {
	var message Message
	message.Title = title
	message.Text = text
	repo.DB.Create(&message)
	return message

}
func (repo *MessageRepository) GetAll() []Message {
	var messages []Message
	repo.DB.Find(&messages)
	return messages

}
