package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"task-tracker-server/internal/domain/task/entity"
	"time"

	"github.com/labstack/echo/v4"
)

type updateTaskRequest struct {
	Title          string     `json:"title"`
	Description    *string    `json:"description"`
	Priority       *int       `json:"priority"`
	StoryPoints    *int       `json:"story_points"`
	TrackedTime    *int       `json:"tracked_time"`
	Deadline       *time.Time `json:"deadline"`
	FinishedAt     *time.Time `json:"finished_at"`
	StatusID       int        `json:"status_id"`
	AssignedUserID *int       `json:"assigned_user_id"`
}

// @Summary	Update an existing task
// @ID			task-update
// @Tags		task
// @Accept		application/json
// @Produce	application/json
// @Param		request	body	updateTaskRequest	true	"Task details to update"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/task/{id} [patch]
func (h *TaskHandler) UpdateTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var request updateTaskRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	err = h.taskUseCase.UpdateTask(c.Request().Context(), &entity.Task{
		ID:             taskID,
		Title:          request.Title,
		Description:    request.Description,
		Priority:       request.Priority,
		StoryPoints:    request.StoryPoints,
		TrackedTime:    request.TrackedTime,
		Deadline:       request.Deadline,
		FinishedAt:     request.FinishedAt,
		StatusID:       request.StatusID,
		AssignedUserID: request.AssignedUserID,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't update task: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
