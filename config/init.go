package config

import (
	"hodei/firstrestapi/modules/message"
	"net/http"
)

func Init() {
	ConnectDB()
	db.Migrator().CreateTable(&message.Message{})
	mRepo := message.MessageRepository{DB: db}
	mServ := message.MessageService{mRepo}
	mController := message.MessageController{mServ}
	RouterInit(mController)
	http.ListenAndServe(":8080", RouterInit(mController))
}
