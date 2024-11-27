package repository

import (
	"context"
	"task-tracker-server/internal/domain/board/entity"
	"task-tracker-server/pkg/postgres"
	"time"
)

type BoardRepository interface {
	CreateBoard(ctx context.Context, board *entity.Board) error
	UpdateBoard(ctx context.Context, board *entity.Board) error
	DeleteBoard(ctx context.Context, boardID int) error
	GetBoard(ctx context.Context, boardID int) (entity.Board, error)
	GetBoardList(ctx context.Context, projectID int) ([]entity.Board, error)
	GetUserBoards(ctx context.Context, userID int) ([]entity.BoardTeam, error)
}

const (
	_defaultConnTimeout = 5 * time.Second
)

type boardRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) BoardRepository {
	return boardRepo{pg}
}
