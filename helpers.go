package main

import (
	"encoding/json" // for encoding and decoding JSON data
	"os"            // command-line arguments and reading/writing files
)

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile) // Open the data file for reading
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(dataFile, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func findTaskByID(tasks []Task, id int) int {
	for i, task := range tasks {
		if task.ID == id {
			return i
		}
	}
	return -1
}

func printUsage() {
	usage := `Usage:
  add <description>   - Add a new task with the given description
  list                - List all tasks
  complete <id>       - Mark the task with the given ID as completed
  delete <id>         - Delete the task with the given ID
  help                - Show this usage message`
	println(usage)
}
