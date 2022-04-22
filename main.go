package main

import (
	"hodei/firstrestapi/config"

	"net/http"
)

func main() {

	config.ConnectDB()

	http.ListenAndServe(":8080", config.RouterInit())
}
