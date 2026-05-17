package entity

import "errors"

// Custom errors for team domain.
var (
	// ErrProjectNotFound is returned when a project is not found.
	ErrProjectNotFound = errors.New("project not found")
	// ErrBoardNotFound is returned when a board is not found.
	ErrBoardNotFound = errors.New("board not found")
	// ErrStatusNotFound is returned when a status is not found.
	ErrStatusNotFound = errors.New("status not found")
	// ErrTaskNotFound is returned when a task is not found.
	ErrTaskNotFound = errors.New("task not found")
)
