package todo

import (
	"encoding/json"
	"os"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Status bool   `json:"status"`
}

var fileName = "todo.json"

func LoadTodo() ([]Todo, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}
	var tasks []Todo
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTodo(tasks []Todo) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
