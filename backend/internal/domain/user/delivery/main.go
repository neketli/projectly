package delivery

import (
	"context"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateAccess(user *entity.User) (string, error)
	CreateRefresh(user *entity.User) (string, error)
	GetUserByRefreshToken(ctx context.Context, requestToken string) (entity.User, error)
}

type UserHandler struct {
	UserUsecase UserUsecase
}

func New(router *echo.Group, usecase UserUsecase) {
	handler := &UserHandler{UserUsecase: usecase}

	h := router.Group("/user")
	{
		h.POST("/register", handler.Register)
		h.POST("/login", handler.Login)
		h.POST("/refresh", handler.Refresh)
	}
}
