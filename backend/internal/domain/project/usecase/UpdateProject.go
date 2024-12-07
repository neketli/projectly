package usecase

import (
	"context"
	"projectly-server/internal/domain/project/entity"
)

func (u *projectUseCase) UpdateProject(ctx context.Context, project *entity.Project) error {
	err := u.repo.UpdateProject(ctx, project)
	if err != nil {
		u.logger.Error("project - usecase - UpdateProject - u.repo.UpdateProject: %s", err.Error())
		return err
	}
	return nil
}
