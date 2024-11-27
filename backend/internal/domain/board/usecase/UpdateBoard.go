package usecase

import (
	"context"
	"task-tracker-server/internal/domain/board/entity"
)

func (u *boardUseCase) UpdateBoard(ctx context.Context, board *entity.Board) error {
	err := u.repo.UpdateBoard(ctx, board)
	if err != nil {
		u.logger.Error("board - usecase - UpdateBoard - u.repo.UpdateBoard: %s", err.Error())
		return err
	}
	return nil
}
