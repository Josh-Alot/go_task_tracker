package tasks

import (
	"encoding/json"
	"log"
	"os"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func CreateTask(tasks []Task) error {
	// get the file path and create a file if it doesn't exist
	filePath := "tasks.json"
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// converts a Go struct into a JSON structure
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(data)
	return err
}
