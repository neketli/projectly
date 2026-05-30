package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/task/entity"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UpdateTask handles updating an existing task.
// @Summary Update an existing task
// @ID			task-update
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		request	body	entity.Task	true	"Task details to update"
// @Success	200	{object}	entity.Task
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/task/{id} [put].
func (h *TaskHandler) UpdateTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var request entity.Task
	if bindErr := c.Bind(&request); bindErr != nil {
		return apierror.Validation("Invalid request body")
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Validation("Failed to authenticate user")
	}

	task, err := h.taskUseCase.UpdateTask(c.Request().Context(), claims.ID, &entity.Task{
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
		return apierror.Internal("Failed to update task")
	}

	return c.JSON(http.StatusOK, task)
}
