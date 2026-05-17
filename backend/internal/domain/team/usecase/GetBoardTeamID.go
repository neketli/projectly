package usecase

import (
	"context"
)

func (u *teamUseCase) GetBoardTeamID(ctx context.Context, boardID int) (int, error) {
	teamID, err := u.repo.GetBoardTeamID(ctx, boardID)
	if err != nil {
		u.logger.Error("team - usecase - GetBoardTeamID - u.repo.GetBoardTeamID: %s", err.Error())
		return 0, err
	}
	return teamID, nil
}
