package usecase

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/user/entity"
)

// GetUserByEmail retrieves a user by email address.
// Получение юзера по email.
func (u *userUseCase) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		u.logger.Error("get user by email failed")
		return entity.User{}, fmt.Errorf("user - usecase - GetUserByEmail: %w", err)
	}
	return user, nil
}
