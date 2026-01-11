package app

import (
	"encoding/json"
	loggo "hard-try/internal/log"
	"log"
	"os"
)

type Questions struct {
	Question string
	Options  []string
	Answer   int
}

func logging() {
	loggo.CreateLog()
}

func readJson() {
	data, err := os.ReadFile("questions.json")
	if err != nil {
		log.Fatal("ошибка:", err)
	}

	var questions []Questions
	err = json.Unmarshal(data, &questions)
	if err != nil {
		log.Fatal("ошибка:", err)
	}
}
