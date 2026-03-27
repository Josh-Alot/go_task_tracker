package main

import (
	"fmt"

	t "github.com/go_task_tracker/tasks"
)

func main() {
	fmt.Printf("Go Task Tracker CLI")

	tasks := []t.Task{
		{Description: "Buy groceries"},
	}

	t.CreateTask(tasks, "tasks.json")

	fmt.Println(t.ListTasks("tasks.json"))
}
