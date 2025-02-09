package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) CreateTask(ctx context.Context, task entity.Task) (entity.TaskDetailed, error) {
	task, err := u.repo.CreateTask(ctx, task)
	if err != nil {
		u.logger.Error("task - usecase - CreateTask - u.repo.CreateTask: %s", err.Error())
		return entity.TaskDetailed{}, err
	}

	detailed, err := u.repo.GetTasks(ctx, &entity.TaskDetailedParams{
		TaskID: &task.ID,
	})
	if err != nil {
		u.logger.Error("task - usecase - CreateTask - u.repo.GetTasks: %s", err.Error())
		return entity.TaskDetailed{}, err
	}

	return detailed[0], nil
}
