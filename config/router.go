package config

import (
	"github.com/gorilla/mux"
	"hodei/firstrestapi/modules/message"
)

func RouterInit(m message.MessageController) *mux.Router {
	router := mux.NewRouter()
	message.RouterInit(db, router)
	return router
}
