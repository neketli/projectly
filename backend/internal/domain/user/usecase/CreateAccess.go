package usecase

import (
	"fmt"
	"projectly-server/internal/domain/user/entity"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func (uc *userUseCase) CreateAccess(user *entity.User) (string, error) {
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(uc.config.AccessTTL))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(uc.config.AccessSecret))
	if err != nil {
		uc.logger.Error(fmt.Errorf("user - usecase - CreateAccess - token.SignedString: %w", err))
		return "", fmt.Errorf("user - usecase - CreateAccess - token.SignedString: %w", err)
	}
	return t, nil
}
