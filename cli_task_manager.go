package main

type Task struct {
	ID          int    `json:"id"`          // Unique identifier for the task
	Description string `json:"description"` // Description of the task
	Completed   bool   `json:"completed"`   // Status of the task (completed or not)
}

const dataFile = "tasks.json" // File to store tasks

func main() {

}
