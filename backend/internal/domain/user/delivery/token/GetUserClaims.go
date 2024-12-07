package token

import (
	"errors"
	"projectly-server/internal/domain/user/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserClaims(c echo.Context) (*entity.JWTClaims, error) {
	user := c.Get("user")
	token, ok := user.(*jwt.Token)
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(*entity.JWTClaims)
	if !ok {
		return nil, errors.New("failed to cast claims as entity.JWTClaims")
	}
	return claims, nil
}
