package usecase

import (
	"context"
)

// DeleteTask deletes a task by ID.
func (u *taskUseCase) DeleteTask(ctx context.Context, taskID int) error {
	err := u.repo.DeleteTask(ctx, taskID)
	if err != nil {
		u.logger.Error("task - usecase - DeleteTask - u.repo.DeleteTask: %s", err.Error())
		return err
	}
	return nil
}
