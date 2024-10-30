package task

import "github.com/google/uuid"

// Repository defines methods for task repository.
type Repository interface {
	// GetTasks will get all task of a user.
	GetTasks(uuid *uuid.UUID) ([]Task, error)

	// AddTask will add a new task to a user.
	AddTask(*Task, *uuid.UUID) (Task, error)

	// UpdateTask will update an existing task.
	UpdateTask(*Task) (Task, error)

	// DeleteTask will delete a task with a specific id.
	DeleteTask(*uuid.UUID) error
}
