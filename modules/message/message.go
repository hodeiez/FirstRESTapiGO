package message

import "gorm.io/gorm"

//TODO: fix dependecy injections, use interface instead
func Init(db *gorm.DB) MessageController {
	mRepo := MessageRepository{DB: db}
	mServ := MessageService{mRepo}
	return MessageController{mServ}
}
