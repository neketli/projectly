package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/task/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Update an existing task
// @ID			task-update
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		request	body	entity.Task	true	"Task details to update"
// @Success	200	{object}	entity.Task
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/task/{id} [put]
func (h *TaskHandler) UpdateTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var request entity.Task
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	task, err := h.taskUseCase.UpdateTask(c.Request().Context(), &entity.Task{
		ID:             taskID,
		Title:          request.Title,
		Description:    request.Description,
		Priority:       request.Priority,
		StoryPoints:    request.StoryPoints,
		TrackedTime:    request.TrackedTime,
		Deadline:       request.Deadline,
		AssignedUserID: request.AssignedUserID,
		FinishedAt:     request.FinishedAt,
		StatusID:       request.StatusID,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't update task: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, task)
}
