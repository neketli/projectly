package usecase

import (
	"context"
)

func (u *boardUseCase) DeleteBoard(ctx context.Context, boardID int) error {
	err := u.repo.DeleteBoard(ctx, boardID)
	if err != nil {
		u.logger.Error("board - usecase - DeleteBoard - u.repo.DeleteBoard: %s", err.Error())
		return err
	}
	return nil
}
