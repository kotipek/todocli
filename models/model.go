package models

import (
	"errors"
	"time"
)

type Status string

const (
	StatusNotStarted Status = "NotStarted"
	StatusInProgress Status = "InProgress"
	StatusCompleted  Status = "Completed"
)

type Task struct {
	ID          string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
}

func ParseStatus(statusInput string) (Status, error) {
	switch statusInput {
	case "NotStarted":
		return StatusNotStarted, nil
	case "InProgress":
		return StatusInProgress, nil
	case "Completed":
		return StatusCompleted, nil
	default:
		return "", errors.New("Неизвестный статус")
	}
}
