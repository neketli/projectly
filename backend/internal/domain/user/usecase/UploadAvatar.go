package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"projectly-server/internal/domain/user/entity"

	"github.com/google/uuid"
)

// UploadAvatar uploads a new avatar for a user.
func (u *userUseCase) UploadAvatar(ctx context.Context, user entity.User, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		u.logger.Error("user - usecase - UploadAvatar - file.Open: %s", err.Error())
		return err
	}

	defer func() {
		if closeErr := src.Close(); closeErr != nil {
			u.logger.Error("user - usecase - UploadAvatar - file.Open: %s", closeErr.Error())
		}
	}()

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
