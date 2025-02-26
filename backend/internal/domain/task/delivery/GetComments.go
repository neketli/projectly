package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/task/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary		Get task comments
// @Description	Get task comments
// @ID				task-get-comments
// @Tags			task
// @Accept			application/json
// @Produce			application/json
// @Param			id				path		string	true	"task_id to get comments"
// @Param			last_comment_id	query		int	true	"id of last comment"
// @Success		200 {object}	[]entity.Comment 	"Array of task comments"
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/task/{id}/comments [get]
func (h *TaskHandler) GetComments(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	lastCommentID, err := strconv.Atoi(c.QueryParam("last_comment_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid comment id",
		}
	}

	var comments []entity.Comment
	comments, err = h.taskUseCase.GetComments(c.Request().Context(), taskID, lastCommentID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get comments: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, comments)
}
