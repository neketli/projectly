package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"
)

func (u *taskUseCase) CreateAttachment(ctx context.Context, taskId int, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		u.logger.Error("user - usecase - UploadAvatar - file.Open: %s", err.Error())
		return "", err
	}

	defer func() {
		if err := src.Close(); err != nil {
			u.logger.Error("task - usecase - CreateAttachment - file.Open: %s", err.Error())
		}
	}()

	filename := fmt.Sprintf("projectly-%d-%s", time.Now().Unix(), file.Filename)

	attachment, err := u.repo.CreateAttachment(ctx, src, filename, taskId)
	if err != nil {
		u.logger.Error("task - usecase - CreateAttachment - u.repo.CreateAttachment: %s", err.Error())
		return "", err
	}

	return attachment, nil
}
