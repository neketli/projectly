package repository

import (
	"context"
	"projectly-server/internal/domain/status/entity"
	"projectly-server/pkg/postgres"
	"time"
)

// StatusRepository defines the interface for status data operations.
type StatusRepository interface {
	CreateStatus(ctx context.Context, status *entity.Status) error
	UpdateStatus(ctx context.Context, status *entity.Status) error
	DeleteStatus(ctx context.Context, statusID int, order int) error
	GetStatusList(ctx context.Context, boardID int) ([]entity.Status, error)
	UpdateOrders(ctx context.Context, boardID int, oldOrder, newOrder int) error
}

const (
	_defaultConnTimeout = 5 * time.Second
)

type statusRepo struct {
	*postgres.Postgres
}

// New creates a new StatusRepository instance.
func New(pg *postgres.Postgres) StatusRepository {
	return statusRepo{pg}
}
