package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// DeleteTask handles the deletion of a task.
// @Summary Delete task
// @ID			task-delete
// @Tags		task
// @Accept		application/json
// @Produce	application/json
// @Param		id	path	int	true	"Task ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/task/{id} [delete].
func (h *TaskHandler) DeleteTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	err = h.taskUseCase.DeleteTask(c.Request().Context(), taskID)
	if err != nil {
		return apierror.Internal("Failed to delete task")
	}

	return c.NoContent(http.StatusOK)
}
