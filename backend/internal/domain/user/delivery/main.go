package delivery

import (
	"context"
	"mime/multipart"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	ChangePassword(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateAccess(user *entity.User) (string, error)
	CreateRefresh(user *entity.User) (string, error)
	GetUserByRefreshToken(ctx context.Context, requestToken string) (entity.User, error)
	UploadAvatar(ctx context.Context, user entity.User, file *multipart.FileHeader) error
	RemoveAvatar(ctx context.Context, user entity.User) error

	CompleteUserAuth(ctx context.Context, user *entity.User) error
}

type UserHandler struct {
	UserUseCase UserUseCase
}

func New(authRouter *echo.Group, router *echo.Group, usecase UserUseCase) {
	handler := &UserHandler{UserUseCase: usecase}

	authRouter.POST("/register", handler.Register)
	authRouter.POST("/login", handler.Login)
	authRouter.POST("/refresh", handler.Refresh)

	authRouter.GET("/:provider", handler.OauthLogin)
	authRouter.GET("/:provider/callback", handler.OauthCallback)

	u := router.Group("/user")
	{
		u.PATCH("/update", handler.Update)
		u.PATCH("/change-password", handler.ChangePassword)
		u.POST("/upload-avatar", handler.UploadAvatar)
		u.DELETE("/remove-avatar", handler.RemoveAvatar)
	}
}
