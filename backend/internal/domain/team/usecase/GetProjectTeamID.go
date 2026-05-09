package usecase

import (
	"context"
)

func (u *teamUseCase) GetProjectTeamID(ctx context.Context, projectID int) (int, error) {
	teamID, err := u.repo.GetProjectTeamID(ctx, projectID)
	if err != nil {
		u.logger.Error("team - usecase - GetProjectTeamID - u.repo.GetProjectTeamID: %s", err.Error())
		return 0, err
	}
	return teamID, nil
}
