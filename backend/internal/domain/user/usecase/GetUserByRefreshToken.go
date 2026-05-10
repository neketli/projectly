package usecase

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/user/entity"

	"github.com/golang-jwt/jwt/v5"
)

// GetUserByRefreshToken retrieves a user by refresh token.
func (u *userUseCase) GetUserByRefreshToken(ctx context.Context, requestToken string) (entity.User, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.config.RefreshSecret), nil
	})
	if err != nil {
		return entity.User{}, fmt.Errorf("user - usecase - GetUserByRefreshToken - jwt.Parse: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return entity.User{}, fmt.Errorf("invalid Token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return entity.User{}, fmt.Errorf("invalid token claims")
	}
	return u.GetUserByEmail(ctx, email)
}
