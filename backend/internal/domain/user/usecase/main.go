package usecase

import (
	"context"
	"mime/multipart"
	"task-tracker-server/config"
	"task-tracker-server/internal/domain/user/entity"
	"task-tracker-server/internal/domain/user/repository"
	"task-tracker-server/pkg/logger"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	ChangePassword(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UploadAvatar(ctx context.Context, user entity.User, file *multipart.FileHeader) error
	RemoveAvatar(ctx context.Context, user entity.User) error
	CreateAccess(user *entity.User) (string, error)
	CreateRefresh(user *entity.User) (string, error)
	GetUserByRefreshToken(ctx context.Context, requestToken string) (entity.User, error)
}

type userUseCase struct {
	repo   repository.UserRepository
	logger *logger.Logger
	config config.Auth
}

func New(r repository.UserRepository, l *logger.Logger, c *config.Config) UserUseCase {
	return &userUseCase{
		repo:   r,
		logger: l,
		config: c.Auth,
	}
}
