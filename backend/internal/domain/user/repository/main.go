package repository

import (
	"context"
	"io"
	"task-tracker-server/internal/domain/user/entity"
	"task-tracker-server/pkg/minio"
	"task-tracker-server/pkg/postgres"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UploadAvatar(ctx context.Context, user entity.User, reader io.Reader, filename string) error
}

const (
	_defaultEntityCap   = 64
	_defaultConnTimeout = 5 * time.Second
)

type userRepo struct {
	*postgres.Postgres
	*minio.Minio
}

func New(pg *postgres.Postgres, s3 *minio.Minio) UserRepository {
	return userRepo{pg, s3}
}
