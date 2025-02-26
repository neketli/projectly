package usecase

import (
	"context"
)

func (u *taskUseCase) CreateComment(ctx context.Context, taskId, userID int, text string) error {
	err := u.repo.CreateComment(ctx, taskId, userID, text)
	if err != nil {
		u.logger.Error("task - usecase - CreateComment - u.repo.CreateComment: %s", err.Error())
		return err
	}

	return nil
}
