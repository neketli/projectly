package usecase

import (
	"context"
	"task-tracker-server/internal/domain/user/entity"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		u.logger.Error("user - usecase - CreateUser - GenerateFromPassword: %s", err.Error())
		return err
	}

	user.Password = string(encryptedPassword)
	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		u.logger.Error("user - usecase - CreateUser - u.repo.CreateUser: %s", err.Error())
		return err
	}
	return nil
}
