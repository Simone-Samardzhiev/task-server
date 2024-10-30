package task

import "github.com/google/uuid"

// Repository defines methods for task repository.
type Repository interface {
	// GetTasks will get all task of a user.
	GetTasks(*uuid.UUID) ([]Task, error)

	// CheckPriority will check if the priority us valid.
	CheckPriority(*int64) (bool, error)

	// AddTask will add a new task to a user.
	AddTask(*Task, *uuid.UUID) error

	// UpdateTask will update an existing task.
	UpdateTask(*Task) error

	// DeleteTask will delete a task with a specific id.
	DeleteTask(*uuid.UUID) error
}
