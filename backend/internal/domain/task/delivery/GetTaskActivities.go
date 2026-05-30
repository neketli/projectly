package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

func (h *TaskHandler) GetTaskActivities(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid task ID")
	}

	activities, err := h.taskUseCase.GetTaskActivities(c.Request().Context(), taskID)
	if err != nil {
		return apierror.Internal("Failed to get task activities")
	}

	return c.JSON(http.StatusOK, activities)
}
