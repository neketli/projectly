package delivery

import (
	"context"
	"mime/multipart"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	ChangePassword(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateAccess(user *entity.User) (string, error)
	CreateRefresh(user *entity.User) (string, error)
	GetUserByRefreshToken(ctx context.Context, requestToken string) (entity.User, error)
	UploadAvatar(ctx context.Context, user entity.User, file *multipart.FileHeader) error
	RemoveAvatar(ctx context.Context, user entity.User) error
}

type UserHandler struct {
	UserUsecase UserUsecase
}

func New(authRouter *echo.Group, router *echo.Group, usecase UserUsecase) {
	handler := &UserHandler{UserUsecase: usecase}

	authRouter.POST("/register", handler.Register)
	authRouter.POST("/login", handler.Login)

	u := router.Group("/user")
	{
		u.POST("/refresh", handler.Refresh)
		u.PATCH("/update", handler.Update)
		u.PATCH("/change-password", handler.ChangePassword)
		u.POST("/upload-avatar", handler.UploadAvatar)
		u.DELETE("/remove-avatar", handler.RemoveAvatar)
	}
}
