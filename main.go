package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HelloObj struct {
	Message string `json:"message"`
}

func GetMessage(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Endpoint Hit: GetMessage")
	json.NewEncoder(w).Encode(HelloObj{"this is a simple JSON message"})
}
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "HELLO REST API")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/message", GetMessage).Methods("GET")
	router.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", router)
}
