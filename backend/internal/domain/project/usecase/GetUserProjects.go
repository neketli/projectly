package usecase

import (
	"context"
	"task-tracker-server/internal/domain/project/entity"
)

// GetUserProjects implements usecase.ProjectUseCase
func (u *projectUseCase) GetUserProjects(ctx context.Context, userID int) ([]entity.Project, error) {
	projects, err := u.repo.GetUserProjects(ctx, userID)
	if err != nil {
		u.logger.Error("project - usecase - GetUserProjects - u.repo.GetUserProjects: %s", err.Error())
		return nil, err
	}

	return projects, nil
}
