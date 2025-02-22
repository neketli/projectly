package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) GetAttachments(ctx context.Context, taskId int) ([]entity.Attachment, error) {
	attachments, err := u.repo.GetAttachments(ctx, taskId)
	if err != nil {
		u.logger.Error("task - usecase - GetAttachments - u.repo.CreateAttachment: %s", err.Error())
		return nil, err
	}

	return attachments, nil
}
