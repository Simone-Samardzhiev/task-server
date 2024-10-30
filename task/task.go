package task

import (
	"github.com/google/uuid"
	"time"
)

// Task defines the data stored in a task.
type Task struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Priority      int64     `json:"priority"`
	DueDate       time.Time `json:"dueDate"`
	DateCompleted NullTime  `json:"dateCompleted"`
	DateDeleted   NullTime  `json:"dateDeleted"`
}

// NewTask is a task that will be added.
type NewTask struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int64     `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
}
