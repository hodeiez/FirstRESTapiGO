package config

import (
	"github.com/gorilla/mux"
	//	"hodei/firstrestapi/modules/hello"
	"hodei/firstrestapi/modules/message"
)

func RouterInit(m message.MessageController) *mux.Router {
	//	router := mux.NewRouter()
	//router.HandleFunc("/message", message.GetMessage).Methods("GET")
	//router.HandleFunc("/", hello.Hello)
	// router.HandleFunc("/create", m.Create).Methods("POST")
	// router.HandleFunc("/messages", m.GetAll).Methods("GET")
	return message.RouterInit(db)
	//return router
}
