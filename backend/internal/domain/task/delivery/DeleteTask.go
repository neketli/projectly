package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete task
// @ID			task-delete
// @Tags		task
// @Accept		application/json
// @Produce	application/json
// @Param		id	path	int	true	"Task ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/task/{id} [delete]
func (h *TaskHandler) DeleteTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	err = h.taskUseCase.DeleteTask(c.Request().Context(), taskID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't delete task: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
