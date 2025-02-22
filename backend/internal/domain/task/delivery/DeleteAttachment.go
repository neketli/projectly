package delivery

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete attachment
// @ID			task-attachment-delete
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		filename	query		string	true	"filename"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/task/delete-attachment [delete]
func (h *TaskHandler) DeleteAttachment(c echo.Context) error {
	filename := c.QueryParam("filename")
	if filename == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid filename",
		}
	}

	err := h.taskUseCase.DeleteAttachment(c.Request().Context(), filename)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't delete attachment: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
