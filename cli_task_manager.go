package main

import (
	"os"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`          // Unique identifier for the task
	Description string `json:"description"` // Description of the task
	Completed   bool   `json:"completed"`   // Status of the task (completed or not)
}

const dataFile = "tasks.json" // File to store tasks

func main() {
	tasks, err := loadTasks() // Load tasks from the file
	if err != nil {
		println("Error loading tasks:", err.Error())
		return
	}

	args := os.Args[1:] // Get command-line arguments (:1 to skip the program name)
	if len(args) == 0 {
		printUsage()
		return
	}

	command := args[0]
	switch command {
	case "add":
		if len(args) < 2 {
			println("Please provide a task description.")
			return
		}
		tasks, err = cmdAdd(tasks, args[1])
		if err != nil {
			println("Error adding task:", err.Error())
			return
		}
		println("Task added successfully.")

	case "list":
		if len(tasks) == 0 {
			println("No tasks found.")
			return
		}
		cmdList(tasks)

	case "complete":
		if len(args) < 2 {
			println("Please provide the ID of the task to complete.")
			return
		}
		taskID, err := strconv.Atoi(args[1])
		if err != nil {
			println("Invalid task ID:", args[1])
			return
		}
		tasks, err = cmdComplete(tasks, taskID)
		if err != nil {
			println("Error completing task:", err.Error())
			return
		}
		println("Task completed successfully.")

	case "edit":
		if len(args) < 3 {
			println("Please provide the ID and Description of the task to edit.")
			return
		}
		taskID, err := strconv.Atoi(args[1])
		if err != nil {
			println("Invalid task ID:", args[1])
			return
		}
		tasks, err = cmdEdit(tasks, taskID, args[2])
		if err != nil {
			println("Error editing task:", err.Error())
			return
		}
		println("Task edited successfully.")

	case "delete":
		if len(args) < 2 {
			println("Please provide the ID of the task to delete.")
			return
		}
		taskID, err := strconv.Atoi(args[1])
		if err != nil {
			println("Invalid task ID:", args[1])
			return
		}
		tasks, err = cmdDelete(tasks, taskID)
		if err != nil {
			println("Error deleting task:", err.Error())
			return
		}
		println("Task deleted successfully.")

	case "help":
		printUsage()

	default:
		println("Unknown command:", command)
		printUsage()
	}

}
