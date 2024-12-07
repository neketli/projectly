package usecase

import (
	"context"
	"projectly-server/internal/domain/team/entity"
)

func (u *teamUseCase) CreateTeam(ctx context.Context, team *entity.Team) error {
	err := u.repo.CreateTeam(ctx, team)
	if err != nil {
		u.logger.Error("team - usecase - CreateTeam - u.repo.CreateTeam: %s", err.Error())
		return err
	}
	return nil
}
