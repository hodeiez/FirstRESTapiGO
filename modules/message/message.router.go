package message

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouterInit(db *gorm.DB, router *mux.Router) {

	controller := Init(db)
	router.HandleFunc("/message", GetMessage).Methods("GET")

	router.HandleFunc("/create", controller.Create).Methods("POST")
	router.HandleFunc("/messages", controller.GetAll).Methods("GET")

}
