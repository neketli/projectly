package usecase

import (
	"context"
	"task-tracker-server/internal/domain/board/entity"
)

// GetUserBoards implements usecase.BoardUseCase
func (u *boardUseCase) GetUserBoards(ctx context.Context, userID int) ([]entity.BoardTeam, error) {
	boards, err := u.repo.GetUserBoards(ctx, userID)
	if err != nil {
		u.logger.Error("board - usecase - GetUserBoards - u.repo.GetUserBoards: %s", err.Error())
		return nil, err
	}

	return boards, nil
}
