package config

import (
	"hodei/firstrestapi/modules/message"
	"net/http"
)

func Init() {
	ConnectDB()
	db.Migrator().CreateTable(&message.Message{})
	mController := message.Init(db)
	http.ListenAndServe(":8080", RouterInit(mController))
}
