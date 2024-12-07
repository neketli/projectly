package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) GetTaskList(ctx context.Context, boardID int) (map[int][]entity.Task, error) {
	tasks, err := u.repo.GetTaskList(ctx, boardID)
	if err != nil {
		u.logger.Error("task - usecase - GetTaskList - u.repo.GetTaskList: %s", err.Error())
		return nil, err
	}

	groupedTasks := make(map[int][]entity.Task)
	for _, task := range tasks {
		groupedTasks[task.StatusID] = append(groupedTasks[task.StatusID], task)
	}
	return groupedTasks, nil
}
