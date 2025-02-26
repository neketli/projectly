package usecase

import (
	"context"
)

func (u *taskUseCase) DeleteComment(ctx context.Context, taskID, commentID int) error {
	err := u.repo.DeleteComment(ctx, taskID, commentID)
	if err != nil {
		u.logger.Error("task - usecase - DeleteComment - u.repo.DeleteComment: %s", err.Error())
		return err
	}

	return nil
}
