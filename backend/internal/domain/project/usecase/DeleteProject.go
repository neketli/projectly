package usecase

import (
	"context"
)

func (u *projectUseCase) DeleteProject(ctx context.Context, projectID int) error {
	err := u.repo.DeleteProject(ctx, projectID)
	if err != nil {
		u.logger.Error("project - usecase - DeleteProject - u.repo.DeleteProject: %s", err.Error())
		return err
	}
	return nil
}
