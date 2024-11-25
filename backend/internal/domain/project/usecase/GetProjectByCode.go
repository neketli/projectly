package usecase

import (
	"context"
	"task-tracker-server/internal/domain/project/entity"
)

func (u *projectUseCase) GetProjectByCode(ctx context.Context, teamID int, code string) (entity.Project, error) {
	project, err := u.repo.GetProjectByCode(ctx, teamID, code)
	if err != nil {
		u.logger.Error("project - usecase - GetProjectByCode - u.repo.GetProjectByCode: %s", err.Error())
		return entity.Project{}, err
	}
	return project, nil
}
