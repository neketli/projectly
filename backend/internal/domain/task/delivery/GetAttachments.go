package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// GetAttachments handles retrieval of task attachments.
// @Summary Get task attachments
// @Description	Get task attachments
// @ID				task-get-attachments
// @Tags			task
// @Accept			application/json
// @Produce			application/json
// @Param			id	path		string	true	"task_id to get attachments"
// @Success		200 {object}	[]string 		"File names"
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/task/{id}/attachments [get].
func (h *TaskHandler) GetAttachments(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	filenames, err := h.taskUseCase.GetAttachments(c.Request().Context(), taskID)
	if err != nil {
		return apierror.Internal("Failed to get attachments")
	}

	return c.JSON(http.StatusOK, filenames)
}
