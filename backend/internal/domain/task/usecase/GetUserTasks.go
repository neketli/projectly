package usecase

import (
	"context"
	"task-tracker-server/internal/domain/task/entity"
)

// GetUserTasks implements usecase.TaskUseCase
func (u *taskUseCase) GetUserTasks(ctx context.Context, userID int, limit uint64) ([]entity.TaskCard, error) {
	tasks, err := u.repo.GetUserTasks(ctx, userID, limit)
	if err != nil {
		u.logger.Error("task - usecase - GetUserTasks - u.repo.GetUserTasks: %s", err.Error())
		return nil, err
	}

	return tasks, nil
}
