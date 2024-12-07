package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	entityTask "projectly-server/internal/domain/task/entity"

	"github.com/labstack/echo/v4"
)

// @Summary	Get task
// @ID			task-get
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		id	path		int	true	"Task id"
// @Success	200	{object}	entityTask.Task
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/task/{id} [get]
func (h *TaskHandler) GetTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var task entityTask.Task
	task, err = h.taskUseCase.GetTask(c.Request().Context(), taskID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get task: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, task)
}
