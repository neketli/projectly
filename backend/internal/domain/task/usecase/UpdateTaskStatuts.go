package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) UpdateTaskStatus(ctx context.Context, task *entity.Task) error {
	err := u.repo.UpdateTaskStatus(ctx, task)
	if err != nil {
		u.logger.Error("task - usecase - UpdateTask - u.repo.UpdateTaskStatus: %s", err.Error())
		return err
	}
	return nil
}
