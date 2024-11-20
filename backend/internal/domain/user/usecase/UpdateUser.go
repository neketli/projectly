package usecase

import (
	"context"
	"task-tracker-server/internal/domain/user/entity"
)

func (u *userUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	err := u.repo.UpdateUser(ctx, user)
	if err != nil {
		u.logger.Error("user - usecase - UpdateUser - u.repo.UpdateUser: %s", err.Error())
		return err
	}
	return nil
}
