package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) GetTask(ctx context.Context, taskID int) (entity.Task, error) {
	task, err := u.repo.GetTask(ctx, taskID)
	if err != nil {
		u.logger.Error("task - usecase - GetTask - u.repo.GetTask: %s", err.Error())
		return entity.Task{}, err
	}
	return task, nil
}
