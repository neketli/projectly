package repository

import (
	"context"
	"task-tracker-server/internal/domain/user/entity"
	"task-tracker-server/pkg/postgres"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

const (
	_defaultEntityCap   = 64
	_defaultConnTimeout = 5 * time.Second
)

type userRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) UserRepository {
	return userRepo{pg}
}
