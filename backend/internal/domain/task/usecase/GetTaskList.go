package usecase

import (
	"context"
	"task-tracker-server/internal/domain/task/entity"
)

func (u *taskUseCase) GetTaskList(ctx context.Context, boardID int, limit uint64) ([]entity.Task, error) {
	tasks, err := u.repo.GetTaskList(ctx, boardID, limit)
	if err != nil {
		u.logger.Error("task - usecase - GetTaskList - u.repo.GetTaskList: %s", err.Error())
		return nil, err
	}
	return tasks, nil
}
