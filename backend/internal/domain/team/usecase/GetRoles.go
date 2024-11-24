package usecase

import (
	"context"
	"task-tracker-server/internal/domain/team/entity"
)

func (u *teamUseCase) GetRoles(ctx context.Context) ([]entity.Role, error) {
	roles, err := u.repo.GetRoles(ctx)
	if err != nil {
		u.logger.Error("team - usecase - GetTeam - u.repo.GetTeam: %s", err.Error())
		return nil, err
	}
	return roles, nil
}
