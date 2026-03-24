package main

import (
	"fmt"

	t "github.com/go_task_tracker/tasks"
)

func main() {
	fmt.Println("Go Task Tracker")

	tasks := []t.Task{
		{ID: 4, Description: "Get daughter from school", Status: t.Status{Name: "todo"}},
	}

	t.CreateTask(tasks)
	fmt.Println(tasks)
}
