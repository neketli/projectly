package delivery

import (
	"net/http"

	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
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

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Validation("Failed to authenticate user")
	}

	err = h.taskUseCase.DeleteAttachment(c.Request().Context(), claims.ID, filename)
	if err != nil {
		return apierror.Internal("Failed to delete attachment")
	}

	return c.NoContent(http.StatusOK)
}
