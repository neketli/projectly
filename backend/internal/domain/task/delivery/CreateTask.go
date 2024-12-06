package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/task/entity"
	"task-tracker-server/internal/domain/user/delivery/token"

	"github.com/labstack/echo/v4"
)

type createTaskRequest struct {
	Title          string  `json:"title"`
	Description    *string `json:"description"`
	Priority       *int    `json:"priority"`
	StoryPoints    *int    `json:"story_points"`
	TrackedTime    *int    `json:"tracked_time"`
	Deadline       *int64  `json:"deadline"`
	FinishedAt     *int64  `json:"finished_at"`
	StatusID       int     `json:"status_id"`
	AssignedUserID *int    `json:"assigned_user_id"`
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

	newTask := entity.Task{
		Title:         request.Title,
		StatusID:      request.StatusID,
		CreatedUserID: claims.ID,
	}
	if request.Description != nil {
		newTask.Description = *request.Description
	}
	if request.Priority != nil {
		newTask.Priority = *request.Priority
	}
	if request.StoryPoints != nil {
		newTask.StoryPoints = *request.StoryPoints
	}
	if request.TrackedTime != nil {
		newTask.TrackedTime = *request.TrackedTime
	}
	if request.Deadline != nil {
		newTask.Deadline = *request.Deadline
	}
	if request.FinishedAt != nil {
		newTask.FinishedAt = *request.FinishedAt
	}
	if request.AssignedUserID != nil {
		newTask.AssignedUserID = *request.AssignedUserID
	}

	task, err := h.taskUseCase.CreateTask(c.Request().Context(), newTask)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't create task: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusCreated, task)
}
