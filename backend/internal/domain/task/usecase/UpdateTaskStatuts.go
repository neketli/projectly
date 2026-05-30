package usecase

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/task/entity"
)

// UpdateTaskStatus updates the status of a task.
func (u *taskUseCase) UpdateTaskStatus(ctx context.Context, userID int, task *entity.Task) error {
	oldTasks, err := u.repo.GetTasks(ctx, &entity.TaskDetailedParams{
		TaskID: &task.ID,
	})
	if err != nil {
		u.logger.Error("task - usecase - UpdateTaskStatus - u.repo.GetTasks: %s", err.Error())
		return err
	}
	if len(oldTasks) == 0 {
		return fmt.Errorf("task not found")
	}
	oldStatusID := oldTasks[0].StatusID

	err = u.repo.UpdateTaskStatus(ctx, task)
	if err != nil {
		u.logger.Error("task - usecase - UpdateTask - u.repo.UpdateTaskStatus: %s", err.Error())
		return err
	}

	u.logActivity(ctx, task.ID, userID, entity.ActionStatusChanged, "status_id",
		fmt.Sprintf("%d", oldStatusID), fmt.Sprintf("%d", task.StatusID))

	return nil
}
