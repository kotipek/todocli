package models

import "time"

type Status string

const (
	StatusNotStarted Status = "NotStarted"
	StatusInProgress Status = "InProgress"
	StatusCompleted  Status = "Completed"
)

type Task struct {
	ID          int
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
}
