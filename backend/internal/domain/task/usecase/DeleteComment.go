package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
)

// DeleteComment deletes a comment from a task.
func (u *taskUseCase) DeleteComment(ctx context.Context, taskID, commentID, userID int) error {
	comment, err := u.repo.GetCommentByID(ctx, commentID)
	if err != nil {
		u.logger.Error("task - usecase - DeleteComment - u.repo.GetCommentByID: %s", err.Error())
		return err
	}

	err = u.repo.DeleteComment(ctx, taskID, commentID)
	if err != nil {
		u.logger.Error("task - usecase - DeleteComment - u.repo.DeleteComment: %s", err.Error())
		return err
	}

	u.logActivity(ctx, taskID, userID, entity.ActionCommentDeleted, "comment", comment.Text, "")

	return nil
}
