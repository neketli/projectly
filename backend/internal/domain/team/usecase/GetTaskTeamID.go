package usecase

import (
	"context"
)

func (u *teamUseCase) GetTaskTeamID(ctx context.Context, taskID int) (int, error) {
	teamID, err := u.repo.GetTaskTeamID(ctx, taskID)
	if err != nil {
		u.logger.Error("team - usecase - GetTaskTeamID - u.repo.GetTaskTeamID: %s", err.Error())
		return 0, err
	}
	return teamID, nil
}
