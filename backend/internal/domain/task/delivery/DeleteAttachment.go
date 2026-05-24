package delivery

import (
	"net/http"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// DeleteAttachment handles deletion of an attachment.
// @Summary Delete attachment
// @ID			task-delete-attachment
// @Tags		task
// @Accept		application/json
// @Produce		application/json
// @Param		filename	query		string	true	"filename"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/task/delete-attachment [delete].
func (h *TaskHandler) DeleteAttachment(c echo.Context) error {
	filename := c.QueryParam("filename")
	if filename == "" {
		return apierror.Validation("Invalid filename")
	}

	err := h.taskUseCase.DeleteAttachment(c.Request().Context(), filename)
	if err != nil {
		return apierror.Internal("Failed to delete attachment")
	}

	return c.NoContent(http.StatusOK)
}
