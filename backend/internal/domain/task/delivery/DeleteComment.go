package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// DeleteComment handles deletion of a comment.
// @Summary Delete task
// @ID			task-delete-comment
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		id			path		int		true	"Task ID"
// @Param		comment_id	query		string	true	"Comment ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/task/{id}/delete-comment [delete].
func (h *TaskHandler) DeleteComment(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid task ID")
	}

	commentID, err := strconv.Atoi(c.QueryParam("comment_id"))
	if commentID == 0 || err != nil {
		return apierror.Validation("Invalid comment ID")
	}

	err = h.taskUseCase.DeleteComment(c.Request().Context(), taskID, commentID)
	if err != nil {
		return apierror.Internal("Failed to delete task")
	}

	return c.NoContent(http.StatusOK)
}
