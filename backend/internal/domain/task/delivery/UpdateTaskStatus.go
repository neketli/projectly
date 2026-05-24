package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/task/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

type requestUpdateTaskStatus struct {
	StatusID   int    `json:"status_id"`
	FinishedAt *int64 `json:"finished_at"`
}

// UpdateTaskStatus handles updating a task status.
// @Summary Update an existing task status
// @ID			task-update-status
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		request	body	requestUpdateTaskStatus	true	"New task status"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/task/{id}/change-status [patch].
func (h *TaskHandler) UpdateTaskStatus(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var request requestUpdateTaskStatus
	if bindErr := c.Bind(&request); bindErr != nil {
		return apierror.Validation("Invalid request body")
	}

	task := entity.Task{
		ID:       taskID,
		StatusID: request.StatusID,
	}

	if request.FinishedAt != nil {
		task.FinishedAt = *request.FinishedAt
	}

	err = h.taskUseCase.UpdateTaskStatus(c.Request().Context(), &task)
	if err != nil {
		return apierror.Internal("Failed to update task")
	}

	return c.NoContent(http.StatusOK)
}
