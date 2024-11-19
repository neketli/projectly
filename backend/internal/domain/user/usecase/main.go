package usecase

import (
	"context"
	"io"
	"mime/multipart"
	"task-tracker-server/config"
	"task-tracker-server/internal/domain/user/entity"
	"task-tracker-server/pkg/logger"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	UploadAvatar(ctx context.Context, user entity.User, reader io.Reader, filename string) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type UserUsecase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User, needHash bool) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UploadAvatar(ctx context.Context, user entity.User, file *multipart.FileHeader) error
	CreateAccess(user *entity.User) (string, error)
	CreateRefresh(user *entity.User) (string, error)
	GetUserByRefreshToken(ctx context.Context, requestToken string) (entity.User, error)
}

type userUseCase struct {
	repo   UserRepository
	logger *logger.Logger
	config config.Auth
}

func New(r UserRepository, l *logger.Logger, c *config.Config) UserUsecase {
	return &userUseCase{
		repo:   r,
		logger: l,
		config: c.Auth,
	}
}
