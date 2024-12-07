package usecase

import (
	"context"
	"projectly-server/internal/domain/board/entity"
)

func (u *boardUseCase) CreateBoard(ctx context.Context, board *entity.Board) error {
	err := u.repo.CreateBoard(ctx, board)
	if err != nil {
		u.logger.Error("board - usecase - CreateBoard - u.repo.CreateBoard: %s", err.Error())
		return err
	}
	return nil
}
