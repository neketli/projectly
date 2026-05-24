package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/task/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetTaskList handles retrieval of task list.
// @Summary Get tasks list
// @Description	Get tasks list of board
// @ID				task-get-list
// @Tags			task
// @Accept			application/json
// @Produce			application/json
// @Param			board_id	query		int	true	"Board id"
// @Success		200			{object}	map[int][]entity.Task
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/task/list [get].
func (h *TaskHandler) GetTaskList(c echo.Context) error {
	boardID, err := strconv.Atoi(c.QueryParam("board_id"))
	if err != nil {
		return apierror.Validation("Invalid board ID")
	}

	var tasks map[int][]entity.Task
	tasks, err = h.taskUseCase.GetTaskList(c.Request().Context(), boardID)
	if err != nil {
		return apierror.Internal("Failed to get tasks")
	}
	return c.JSON(http.StatusOK, tasks)
}
