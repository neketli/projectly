package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) CreateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	task, err := u.repo.CreateTask(ctx, task)
	if err != nil {
		u.logger.Error("task - usecase - CreateTask - u.repo.CreateTask: %s", err.Error())
		return entity.Task{}, err
	}
	return task, nil
}
