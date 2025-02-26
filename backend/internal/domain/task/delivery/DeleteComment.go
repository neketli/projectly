package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete task
// @ID			task-delete-comment
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		id			path		int		true	"Task ID"
// @Param		comment_id	query		string	true	"Comment ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/task/{id}/delete-comment [delete]
func (h *TaskHandler) DeleteComment(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid task id",
		}
	}

	commentID, err := strconv.Atoi(c.QueryParam("comment_id"))
	if commentID == 0 || err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid comment_id",
		}
	}

	err = h.taskUseCase.DeleteComment(c.Request().Context(), taskID, commentID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't delete task: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
