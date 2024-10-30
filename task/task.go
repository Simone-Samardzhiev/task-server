package task

import (
	"time"
)

type Task struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Priority      int       `json:"priority"`
	DueDate       time.Time `json:"dueDate"`
	DateCompleted NullTime  `json:"dateCompleted"`
	DateDeleted   NullTime  `json:"dateDeleted"`
}
