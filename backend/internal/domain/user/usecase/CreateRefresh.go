package usecase

import (
	"fmt"
	"strconv"
	"task-tracker-server/internal/domain/user/entity"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type refreshCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (uc *userUseCase) CreateRefresh(user *entity.User) (string, error) {
	claims := &refreshCustomClaims{
		ID:    strconv.Itoa(user.ID),
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(uc.config.RefreshTTL))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte(uc.config.RefreshSecret))
	if err != nil {
		uc.logger.Error(fmt.Errorf("user - usecase - CreateRefresh - token.SignedString: %w", err))
		return "", fmt.Errorf("user - usecase - CreateRefresh - token.SignedString: %w", err)
	}
	return rt, nil
}
