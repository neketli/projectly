package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

// UpdateTask updates a task and returns its detailed representation.
func (u *taskUseCase) UpdateTask(ctx context.Context, task *entity.Task) (entity.TaskDetailed, error) {
	err := u.repo.UpdateTask(ctx, task)
	if err != nil {
		u.logger.Error("task - usecase - UpdateTask - u.repo.UpdateTask: %s", err.Error())
		return entity.TaskDetailed{}, err
	}

	detailed, err := u.repo.GetTasks(ctx, &entity.TaskDetailedParams{
		TaskID: &task.ID,
	})
	if err != nil {
		u.logger.Error("task - usecase - UpdateTask - u.repo.GetTasks: %s", err.Error())
		return entity.TaskDetailed{}, err
	}

	return detailed[0], nil
}
