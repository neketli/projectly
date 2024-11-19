package usecase

import (
	"context"
	"task-tracker-server/internal/domain/user/entity"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUseCase) UpdateUser(ctx context.Context, user *entity.User, needHash bool) error {
	if needHash {
		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			u.logger.Error("user - usecase - UpdateUser - GenerateFromPassword: %s", err.Error())
			return err
		}

		user.Password = string(encryptedPassword)
	}

	err := u.repo.UpdateUser(ctx, user)
	if err != nil {
		u.logger.Error("user - usecase - UpdateUser - u.repo.UpdateUser: %s", err.Error())
		return err
	}
	return nil
}
