package main

import (
	"fmt"

	t "github.com/go_task_tracker/tasks"
)

func main() {
	fmt.Println("Go Task Tracker")

	tasks := []t.Task{
		{ID: 1, Description: "Buy groceries", Status: "todo"},
	}

	t.CreateTask(tasks)
	fmt.Println(tasks)
}
