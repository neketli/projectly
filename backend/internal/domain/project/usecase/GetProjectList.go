package usecase

import (
	"context"
	"projectly-server/internal/domain/project/entity"
)

func (u *projectUseCase) GetProjectList(ctx context.Context, teamID int) ([]entity.Project, error) {
	projects, err := u.repo.GetProjectList(ctx, teamID)
	if err != nil {
		u.logger.Error("project - usecase - GetProjectList - u.repo.GetProjectList: %s", err.Error())
		return nil, err
	}
	return projects, nil
}
