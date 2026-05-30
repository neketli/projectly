package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) GetTaskActivities(ctx context.Context, taskID int) ([]entity.TaskActivity, error) {
	activities, err := u.repo.GetTaskActivities(ctx, taskID)
	if err != nil {
		u.logger.Error("task - usecase - GetTaskActivities - u.repo.GetTaskActivities: %s", err.Error())
		return nil, err
	}

	return activities, nil
}
