package usecase

import (
	"context"
	"projectly-server/internal/domain/board/entity"
)

func (u *boardUseCase) GetBoardList(ctx context.Context, projectID int) ([]entity.Board, error) {
	boards, err := u.repo.GetBoardList(ctx, projectID)
	if err != nil {
		u.logger.Error("board - usecase - GetBoardList - u.repo.GetBoardList: %s", err.Error())
		return nil, err
	}
	return boards, nil
}
