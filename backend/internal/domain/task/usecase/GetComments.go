package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

func (u *taskUseCase) GetComments(ctx context.Context, taskId, lastCommentID int) ([]entity.Comment, error) {
	comments, err := u.repo.GetComments(ctx, taskId, lastCommentID)
	if err != nil {
		u.logger.Error("task - usecase - GetComments - u.repo.GetComments: %s", err.Error())
		return nil, err
	}

	return comments, nil
}
