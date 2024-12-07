package usecase

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/user/entity"
)

// Получение юзера по email
func (uc *userUseCase) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		uc.logger.Error("get user by email failed")
		return entity.User{}, fmt.Errorf("user - usecase - GetUserByEmail: %w", err)
	}
	return user, nil
}
