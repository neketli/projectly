package usecase

import (
	"context"
	"projectly-server/internal/domain/user/entity"
)

func (u *userUseCase) RemoveAvatar(ctx context.Context, user entity.User) error {
	err := u.repo.RemoveAvatar(ctx, user.ID, user.Meta.Avatar)
	if err != nil {
		u.logger.Error("user - usecase - UploadAvatar - u.repo.RemoveAvatar: %s", err.Error())
		return err
	}

	return nil
}
