package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/task/entity"
	"task-tracker-server/internal/domain/user/delivery/token"
	"time"

	"github.com/labstack/echo/v4"
)

type createTaskRequest struct {
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

// @Summary	Create a new task
// @ID			task-create
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		request	body		createTaskRequest	true	"New task details"
// @Success	201		{object}	entity.Task			"Created task"
// @Failure	400		{object}	echo.HTTPError		"Bad request"
// @Failure	500		{object}	echo.HTTPError		"Internal server error"
// @Router		/task/create [post]
func (h *TaskHandler) CreateTask(c echo.Context) error {
	var request createTaskRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	task, err := h.taskUseCase.CreateTask(c.Request().Context(), entity.Task{
		Title:          request.Title,
		Description:    request.Description,
		Priority:       request.Priority,
		StoryPoints:    request.StoryPoints,
		TrackedTime:    request.TrackedTime,
		Deadline:       request.Deadline,
		FinishedAt:     request.FinishedAt,
		StatusID:       request.StatusID,
		CreatedUserID:  claims.ID,
		AssignedUserID: request.AssignedUserID,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't create task: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusCreated, task)
}
