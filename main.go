package main

import (
	"hodei/firstrestapi/config"
)

func main() {

	//config.ConnectDB()
	config.Init()
	//http.ListenAndServe(":8080", config.RouterInit())
}
