package storage

import (
	"encoding/json"
	"os"
)

type Questions struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
}

func ReadQuestions() ([]Questions, error) {
	data, err := os.ReadFile("questions.json")
	if err != nil {
		return nil, err
	}

	var questions []Questions
	err = json.Unmarshal(data, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}
