package entity

import "errors"

var (
	ErrProjectNotFound = errors.New("project not found")
	ErrBoardNotFound   = errors.New("board not found")
	ErrStatusNotFound  = errors.New("status not found")
	ErrTaskNotFound     = errors.New("task not found")
)
