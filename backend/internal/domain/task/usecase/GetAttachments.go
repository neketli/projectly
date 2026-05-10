package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

// GetAttachments retrieves all attachments for a task.
func (u *taskUseCase) GetAttachments(ctx context.Context, taskID int) ([]entity.Attachment, error) {
	attachments, err := u.repo.GetAttachments(ctx, taskID)
	if err != nil {
		u.logger.Error("task - usecase - GetAttachments - u.repo.CreateAttachment: %s", err.Error())
		return nil, err
	}

	return attachments, nil
}
