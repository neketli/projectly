package usecase

import (
	"fmt"
	"projectly-server/internal/domain/user/entity"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// refreshCustomClaims represents JWT claims for refresh token.
type refreshCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// CreateRefresh generates a refresh token for a user.
func (u *userUseCase) CreateRefresh(user *entity.User) (string, error) {
	claims := &refreshCustomClaims{
		ID:    strconv.Itoa(user.ID),
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(u.config.RefreshTTL))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte(u.config.RefreshSecret))
	if err != nil {
		u.logger.Error(fmt.Errorf("user - usecase - CreateRefresh - token.SignedString: %w", err))
		return "", fmt.Errorf("user - usecase - CreateRefresh - token.SignedString: %w", err)
	}
	return rt, nil
}
