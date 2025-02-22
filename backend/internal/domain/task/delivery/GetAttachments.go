package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary		Get task attachments
// @Description	Get task attachments
// @ID				get-task-attachments
// @Tags			task
// @Accept			application/json
// @Produce			application/json
// @Param			id	path		string	true	"task_id to get attachments"
// @Success		200 {object}	[]string 		"File names"
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/task/{id}/attachments [get]
func (h *TaskHandler) GetAttachments(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	filenames, err := h.taskUseCase.GetAttachments(c.Request().Context(), taskID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get attachments: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, filenames)
}
