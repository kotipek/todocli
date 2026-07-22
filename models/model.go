package models

import "time"

type Status string

const (
	StatusNotStarted Status = "NotStarted"
	StatusInProgress Status = "InProgress"
	StatusCompleted  Status = "Completed"
)

type Task struct {
	Name        string
	Status      Status
	Date        time.Time
	Description string
}
