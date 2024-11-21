package usecase

import (
	"context"
	"task-tracker-server/internal/domain/team/entity"
)

func (u *teamUseCase) UpdateTeam(ctx context.Context, team *entity.Team) error {
	err := u.repo.UpdateTeam(ctx, team)
	if err != nil {
		u.logger.Error("team - usecase - UpdateTeam - u.repo.UpdateTeam: %s", err.Error())
		return err
	}
	return nil
}
