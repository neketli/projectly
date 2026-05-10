package usecase

import (
	"context"
)

// CreateComment creates a comment for a task.
func (u *taskUseCase) CreateComment(ctx context.Context, taskID, userID int, text string) error {
	err := u.repo.CreateComment(ctx, taskID, userID, text)
	if err != nil {
		u.logger.Error("task - usecase - CreateComment - u.repo.CreateComment: %s", err.Error())
		return err
	}

	return nil
}
