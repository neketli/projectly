package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) UpdateTask(ctx context.Context, task *entity.Task) error {
	err := u.repo.UpdateTask(ctx, task)
	if err != nil {
		u.logger.Error("task - usecase - UpdateTask - u.repo.UpdateTask: %s", err.Error())
		return err
	}
	return nil
}
