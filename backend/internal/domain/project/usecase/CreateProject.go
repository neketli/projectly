package usecase

import (
	"context"
	"projectly-server/internal/domain/project/entity"
)

func (u *projectUseCase) CreateProject(ctx context.Context, project *entity.Project) error {
	err := u.repo.CreateProject(ctx, project)
	if err != nil {
		u.logger.Error("project - usecase - CreateProject - u.repo.CreateProject: %s", err.Error())
		return err
	}
	return nil
}
