package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) GetTasks(ctx context.Context, params *entity.TaskDetailedParams) ([]entity.TaskDetailed, error) {
	tasks, err := u.repo.GetTasks(ctx, params)
	if err != nil {
		u.logger.Error("task - usecase - GetTasks - u.repo.GetTasks: %s", err.Error())
		return nil, err
	}

	return tasks, nil
}
