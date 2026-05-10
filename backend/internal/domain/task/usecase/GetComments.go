package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

// GetComments retrieves comments for a task.
func (u *taskUseCase) GetComments(ctx context.Context, taskID, lastCommentID int) ([]entity.Comment, error) {
	comments, err := u.repo.GetComments(ctx, taskID, lastCommentID)
	if err != nil {
		u.logger.Error("task - usecase - GetComments - u.repo.GetComments: %s", err.Error())
		return nil, err
	}

	return comments, nil
}
