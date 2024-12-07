package usecase

import (
	"context"
	"projectly-server/internal/domain/project/entity"
)

func (u *projectUseCase) GetProject(ctx context.Context, projectID int) (entity.Project, error) {
	project, err := u.repo.GetProject(ctx, projectID)
	if err != nil {
		u.logger.Error("project - usecase - GetProject - u.repo.GetProject: %s", err.Error())
		return entity.Project{}, err
	}
	return project, nil
}
