package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	entityTask "projectly-server/internal/domain/task/entity"

	"github.com/labstack/echo/v4"
)

// GetTask handles retrieval of a task.
// @Summary Get task
// @ID			task-get
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		id	path		int	true	"Task id"
// @Success	200	{object}	entityTask.TaskDetailed
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/task/{id} [get].
func (h *TaskHandler) GetTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var tasks []entityTask.TaskDetailed
	tasks, err = h.taskUseCase.GetTasks(c.Request().Context(), &entityTask.TaskDetailedParams{
		TaskID: &taskID,
	})
	if err != nil || tasks == nil || len(tasks) == 0 {
		return apierror.Internal("Failed to get task")
	}

	return c.JSON(http.StatusOK, tasks[0])
}
