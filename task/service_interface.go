package task

import "github.com/google/uuid"

// Service defines methods for task service.
type Service interface {
	// GetTasks will return all tasks of a user.
	GetTasks(*string) ([]Task, error)

	// AddTask will add a new task to a user.
	AddTask(*string, *NewTask) (*Task, error)

	// UpdateTask will update an existing task.
	UpdateTask(*string, *Task) (*Task, error)

	// DeleteTask will delete an existing task.
	DeleteTask(*string, *uuid.UUID) error
}
