package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// DeleteStatus handles the deletion of a status.
// @Summary Delete status
// @ID			status-delete
// @Tags		status
// @Accept		application/json
// @Produce		application/json
// @Param		id		query		int	true	"Status id"
// @Param		order	query		int	true	"Status order"
// @Success	204
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/status/delete [delete].
func (h *StatusHandler) DeleteStatus(c echo.Context) error {
	statusID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return apierror.Validation("Invalid status ID")
	}

	order, err := strconv.Atoi(c.QueryParam("order"))
	if err != nil {
		return apierror.Validation("Invalid status order")
	}

	err = h.statusUseCase.DeleteStatus(c.Request().Context(), statusID, order)
	if err != nil {
		return apierror.Internal("Failed to delete status")
	}

	return c.NoContent(http.StatusNoContent)
}
