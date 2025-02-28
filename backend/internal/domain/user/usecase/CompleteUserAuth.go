package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"projectly-server/internal/domain/user/entity"
)

func (u *userUseCase) CompleteUserAuth(ctx context.Context, user *entity.User) error {
	existingUser, err := u.GetUserByEmail(ctx, user.Email)
	if err != nil && err != entity.ErrNoUserFound {
		return fmt.Errorf("failed to check existing user info: %v", err)
	}
	if err == entity.ErrNoUserFound {
		user.Password = generatePassword(16)

		err = u.CreateUser(ctx, user)
		if err != nil {
			return fmt.Errorf("failed to create user: %v", err)
		}
		return nil
	}

	if existingUser.Meta.Provider == user.Meta.Provider {
		return nil
	}

	existingUser.Meta = &entity.UserMeta{
		Provider:   user.Meta.Provider,
		ProviderID: user.Meta.ProviderID,
	}
	err = u.UpdateUser(ctx, &existingUser)
	if err != nil {
		return fmt.Errorf("failed to update user info: %v", err)
	}
	return nil

}

func generatePassword(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}
