package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"projectly-server/internal/domain/task/entity"
	"time"
)

// CreateAttachment uploads an attachment for a task.
func (u *taskUseCase) CreateAttachment(ctx context.Context, taskID, userID int, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		u.logger.Error("user - usecase - UploadAvatar - file.Open: %s", err.Error())
		return "", err
	}

	defer func() {
		if closeErr := src.Close(); closeErr != nil {
			u.logger.Error("task - usecase - CreateAttachment - file.Open: %s", closeErr.Error())
		}
	}()

	filename := fmt.Sprintf("projectly-%d-%s", time.Now().Unix(), file.Filename)

	attachment, err := u.repo.CreateAttachment(ctx, src, filename, taskID)
	if err != nil {
		u.logger.Error("task - usecase - CreateAttachment - u.repo.CreateAttachment: %s", err.Error())
		return "", err
	}

	u.logActivity(ctx, taskID, userID, entity.ActionAttachmentAdded, "attachment", "", file.Filename)

	return attachment, nil
}
