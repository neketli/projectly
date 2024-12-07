package usecase

import (
	"context"
	"projectly-server/internal/domain/board/entity"
	"projectly-server/internal/domain/board/repository"
	"projectly-server/pkg/logger"
)

type BoardUseCase interface {
	CreateBoard(ctx context.Context, board *entity.Board) error
	UpdateBoard(ctx context.Context, board *entity.Board) error
	DeleteBoard(ctx context.Context, boardID int) error
	GetBoard(ctx context.Context, boardID int) (entity.Board, error)
	GetBoardList(ctx context.Context, projectID int) ([]entity.Board, error)
	GetUserBoards(ctx context.Context, userID int) ([]entity.BoardTeam, error)
}

type boardUseCase struct {
	repo   repository.BoardRepository
	logger *logger.Logger
}

func New(r repository.BoardRepository, l *logger.Logger) BoardUseCase {
	return &boardUseCase{
		repo:   r,
		logger: l,
	}
}
