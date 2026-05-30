package usecase

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/task/entity"
)

// DeleteAttachment deletes an attachment by filename.
func (u *taskUseCase) DeleteAttachment(ctx context.Context, userID int, objectName string) error {
	taskID, err := u.parseTaskIDFromObjectName(objectName)
	if err != nil {
		u.logger.Error("task - usecase - DeleteAttachment - parseTaskIDFromObjectName: %s", err.Error())
	}

	err = u.repo.DeleteAttachment(ctx, objectName)
	if err != nil {
		u.logger.Error("task - usecase - DeleteAttachment - u.repo.DeleteAttachment: %s", err.Error())
		return err
	}

	if taskID > 0 {
		u.logActivity(ctx, taskID, userID, entity.ActionAttachmentDeleted, "attachment", objectName, "")
	}

	return nil
}

func (u *taskUseCase) parseTaskIDFromObjectName(objectName string) (int, error) {
	var taskID int
	_, err := fmt.Sscanf(objectName, "attachments/%d/", &taskID)
	return taskID, err
}
