package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/google/uuid"
)

func (u *userUseCase) UploadAvatar(ctx context.Context, user entity.User, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		u.logger.Error("user - usecase - UploadAvatar - file.Open: %s", err.Error())
		return err
	}

	defer src.Close()

	if user.Meta != nil && user.Meta.Avatar != "" {
		err = u.repo.RemoveAvatar(ctx, user.ID, user.Meta.Avatar)
		if err != nil {
			u.logger.Error("user - usecase - UploadAvatar - u.repo.RemoveAvatar: %s", err.Error())
			return err
		}
	}

	filename := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))
	err = u.repo.UploadAvatar(ctx, user, src, filename)
	if err != nil {
		u.logger.Error("user - usecase - UploadAvatar - u.repo.UploadAvatar: %s", err.Error())
		return err
	}
	return nil
}
