package delivery

import (
	"fmt"
	"net/http"

	entityTask "projectly-server/internal/domain/task/entity"

	"github.com/labstack/echo/v4"
)

// @Summary		Get detailed tasks list
// @ID			task-get-detailed
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param			limit			query		int		false	"Task limits"
// @Param			user_id			query		int		false	"User id"
// @Param			task_id			query		int		false	"Task id"
// @Param			team_id			query		int		false	"Team id"
// @Param			board_id		query		int		false	"Board id"
// @Param			project_code	query		int		false	"Project code"
// @Param			project_index	query		int		false	"Task index in project"
// @Param			search			query		string	false	"Search text in tasks"
// @Success	200		{array}		entityTask.TaskDetailed
// @Failure	400		{object}	echo.HTTPError
// @Failure	500		{object}	echo.HTTPError
// @Router		/task/ [get]
func (h *TaskHandler) GetTasks(c echo.Context) error {
	var params entityTask.TaskDetailedParams
	err := c.Bind(&params)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %s", err.Error()),
		}
	}

	var tasks []entityTask.TaskDetailed
	tasks, err = h.taskUseCase.GetTasks(c.Request().Context(), &params)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get tasks: %s", err.Error()),
		}
	}
	return c.JSON(http.StatusOK, tasks)
}
