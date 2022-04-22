package message

import (
	//"time"

	"gorm.io/gorm"
)

type HelloObj struct {
	Message string `json:"message"`
}
type Message struct {
	gorm.Model
	Title string
	Text  string
}
