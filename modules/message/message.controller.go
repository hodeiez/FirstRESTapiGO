package message

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MessageController struct {
	MessageService
}

func GetMessage(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Endpoint Hit: GetMessage")
	json.NewEncoder(w).Encode(HelloObj{"this is a simple JSON message"})
}
func (c *MessageController) Create(w http.ResponseWriter, req *http.Request) {

	json.NewEncoder(w).Encode(c.MessageService.Create(req.Body))

}
func (c *MessageController) GetAll(w http.ResponseWriter, req *http.Request) {

	json.NewEncoder(w).Encode(c.MessageService.GetAll())

}
