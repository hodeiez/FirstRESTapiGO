package message

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMessage(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Endpoint Hit: GetMessage")
	json.NewEncoder(w).Encode(HelloObj{"this is a simple JSON message"})
}
