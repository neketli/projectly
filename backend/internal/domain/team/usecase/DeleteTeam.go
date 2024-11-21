package usecase

import (
	"context"
)

func (u *teamUseCase) DeleteTeam(ctx context.Context, teamID int) error {
	err := u.repo.DeleteTeam(ctx, teamID)
	if err != nil {
		u.logger.Error("team - usecase - DeleteTeam - u.repo.DeleteTeam: %s", err.Error())
		return err
	}
	return nil
}
