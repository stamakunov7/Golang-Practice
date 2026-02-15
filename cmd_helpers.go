package main

import (
	"os"
)

func cmdAdd(tasks []Task, description string) ([]Task, error) {
	newTask := Task{
		ID:          nextID(tasks),
		Description: description,
		Completed:   false,
	}
	tasks = append(tasks, newTask)
	err := saveTasks(tasks)
	return tasks, err
}

func cmdList(tasks []Task) {
	if len(tasks) == 0 {
		println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		println("[", status, "]", task.ID, "-", task.Description)
	}
}

func cmdComplete(tasks []Task, id int) ([]Task, error) {
	index := findTaskByID(tasks, id)
	if index == -1 {
		return tasks, os.ErrNotExist
	}
	tasks[index].Completed = true
	err := saveTasks(tasks)
	return tasks, err
}

func cmdDelete(tasks []Task, id int) ([]Task, error) {
	index := findTaskByID(tasks, id)
	if index == -1 {
		return tasks, os.ErrNotExist
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	err := saveTasks(tasks)
	return tasks, err
}
