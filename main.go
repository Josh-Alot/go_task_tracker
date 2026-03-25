package main

import (
	"fmt"

	t "github.com/go_task_tracker/tasks"
)

func main() {
	tasks, _ := t.ListTasks("tasks.json")

	fmt.Println(tasks)
}
