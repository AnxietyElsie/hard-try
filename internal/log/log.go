package loggo

import (
	"log"
	"os"
)

func CreateLog() {
	logCreate, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("ошибка:", err)
	}
	defer logCreate.Close()
}
