package config

import (
	"github.com/gorilla/mux"
	"hodei/firstrestapi/modules/hello"
	"hodei/firstrestapi/modules/message"
)

func RouterInit() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/message", message.GetMessage).Methods("GET")
	router.HandleFunc("/", hello.Hello)
	return router
}
