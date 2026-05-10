package usecase

import (
	"fmt"
	"projectly-server/internal/domain/user/entity"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CreateAccess generates an access token for a user.
func (u *userUseCase) CreateAccess(user *entity.User) (string, error) {
	avatar := ""
	if user.Meta != nil {
		avatar = user.Meta.Avatar
	}
	claims := &entity.JWTClaims{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Meta: entity.UserMeta{
			Avatar: avatar,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(u.config.AccessTTL))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(u.config.AccessSecret))
	if err != nil {
		u.logger.Error(fmt.Errorf("user - usecase - CreateAccess - token.SignedString: %w", err))
		return "", fmt.Errorf("user - usecase - CreateAccess - token.SignedString: %w", err)
	}
	return t, nil
}
