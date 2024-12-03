package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	entityTask "task-tracker-server/internal/domain/task/entity"

	"github.com/labstack/echo/v4"
)

// @Summary		Get tasks list
// @Description	Get tasks list of board
// @ID				task-get-list
// @Tags			task
// @Accept			application/json
// @Produce		application/json
// @Param			limit	query		int	false	"Task limits"
// @Param			board_id	query		int	true	"Board id"
// @Success		200			{array}		entityTask.Task
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/task/list [get]
func (h *TaskHandler) GetTaskList(c echo.Context) error {
	boardID, err := strconv.Atoi(c.QueryParam("board_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid board id",
		}
	}

	var limit uint64
	queryLimit := c.QueryParam("limit")
	if queryLimit != "" {
		l, err := strconv.Atoi(queryLimit)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "invalid board id",
			}
		}
		limit = uint64(l)
	}

	var tasks []entityTask.Task
	tasks, err = h.taskUseCase.GetTaskList(c.Request().Context(), boardID, limit)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get tasks: %s", err.Error()),
		}
	}
	return c.JSON(http.StatusOK, tasks)
}
