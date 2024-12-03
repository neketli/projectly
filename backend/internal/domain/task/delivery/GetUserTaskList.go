package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	entityTask "task-tracker-server/internal/domain/task/entity"

	"github.com/labstack/echo/v4"
)

// @Summary	Get users tasks list
// @ID			task-user-get-list
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param			limit	query		int	false	"Task limits"
// @Param			user_id	query		int	true	"User id"
// @Success	200		{array}		entityTask.TaskCard
// @Failure	400		{object}	echo.HTTPError
// @Failure	500		{object}	echo.HTTPError
// @Router		/task/list-user [get]
func (h *TaskHandler) GetUserTaskList(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
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

	var tasks []entityTask.TaskCard
	tasks, err = h.taskUseCase.GetUserTasks(c.Request().Context(), userID, limit)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get tasks: %s", err.Error()),
		}
	}
	return c.JSON(http.StatusOK, tasks)
}
