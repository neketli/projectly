package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"projectly-server/internal/domain/task/entity"

	"github.com/labstack/echo/v4"
)

type requestUpdateTaskStatus struct {
	StatusID int `json:"status_id"`
}

// @Summary	Update an existing task status
// @ID			task-update-status
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		request	body	requestUpdateTaskStatus	true	"New task status"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/task/{id}/change-status [patch]
func (h *TaskHandler) UpdateTaskStatus(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var request requestUpdateTaskStatus
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	err = h.taskUseCase.UpdateTaskStatus(c.Request().Context(), &entity.Task{
		ID:       taskID,
		StatusID: request.StatusID,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't update task: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
