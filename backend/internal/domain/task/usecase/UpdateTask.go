package usecase

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/task/entity"
)

// UpdateTask updates a task and returns its detailed representation.
func (u *taskUseCase) UpdateTask(ctx context.Context, userID int, task *entity.Task) (entity.TaskDetailed, error) {
	oldTasks, err := u.repo.GetTasks(ctx, &entity.TaskDetailedParams{
		TaskID: &task.ID,
	})
	if err != nil {
		u.logger.Error("task - usecase - UpdateTask - u.repo.GetTasks: %s", err.Error())
		return entity.TaskDetailed{}, err
	}
	if len(oldTasks) == 0 {
		return entity.TaskDetailed{}, fmt.Errorf("task not found")
	}
	oldTask := oldTasks[0]

	err = u.repo.UpdateTask(ctx, task)
	if err != nil {
		u.logger.Error("task - usecase - UpdateTask - u.repo.UpdateTask: %s", err.Error())
		return entity.TaskDetailed{}, err
	}

	if oldTask.Title != task.Title {
		u.logActivity(ctx, task.ID, userID, entity.ActionTaskUpdated, "title", oldTask.Title, task.Title)
	}
	if oldTask.Description != task.Description {
		u.logActivity(ctx, task.ID, userID, entity.ActionTaskUpdated, "description", oldTask.Description, task.Description)
	}
	if oldTask.Priority != task.Priority {
		u.logActivity(ctx, task.ID, userID, entity.ActionTaskUpdated, "priority", fmt.Sprintf("%d", oldTask.Priority), fmt.Sprintf("%d", task.Priority))
	}
	if oldTask.AssignedUserID != task.AssignedUserID {
		oldAssigned := fmt.Sprintf("%d", oldTask.AssignedUserID)
		newAssigned := fmt.Sprintf("%d", task.AssignedUserID)
		u.logActivity(ctx, task.ID, userID, entity.ActionTaskUpdated, "assigned_user_id", oldAssigned, newAssigned)
	}
	if oldTask.Deadline != task.Deadline {
		oldDeadline := fmt.Sprintf("%d", oldTask.Deadline)
		newDeadline := fmt.Sprintf("%d", task.Deadline)
		u.logActivity(ctx, task.ID, userID, entity.ActionTaskUpdated, "deadline", oldDeadline, newDeadline)
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

func (u *taskUseCase) logActivity(ctx context.Context, taskID, userID int, actionType, fieldName, oldValue, newValue string) {
	err := u.repo.CreateTaskActivity(ctx, taskID, userID, actionType, fieldName, oldValue, newValue)
	if err != nil {
		u.logger.Error("task - usecase - logActivity - u.repo.CreateTaskActivity: %s", err.Error())
	}
}
