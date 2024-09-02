package test_service

import "github.com/google/uuid"

type Task struct {
	ID string
}

func (t Task) TaskID() string {
	return t.ID
}

func GetTasks() []Task {
	n := 100
	tasks := make([]Task, n)

	for i := range tasks {
		tasks[i].ID = uuid.New().String()
	}

	return tasks
}
