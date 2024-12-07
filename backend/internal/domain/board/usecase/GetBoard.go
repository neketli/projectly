package usecase

import (
	"context"
	"projectly-server/internal/domain/board/entity"
)

func (u *boardUseCase) GetBoard(ctx context.Context, boardID int) (entity.Board, error) {
	board, err := u.repo.GetBoard(ctx, boardID)
	if err != nil {
		u.logger.Error("board - usecase - GetBoard - u.repo.GetBoard: %s", err.Error())
		return entity.Board{}, err
	}
	return board, nil
}
