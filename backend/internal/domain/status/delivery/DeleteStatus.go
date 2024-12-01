package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete status
// @ID			status-delete
// @Tags		status
// @Accept		application/json
// @Produce		application/json
// @Param		id		query		int	true	"Status id"
// @Param		order	query		int	true	"Status order"
// @Success	204
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/status/delete [delete]
func (h *StatusHandler) DeleteStatus(c echo.Context) error {
	statusID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid status id",
		}
	}

	order, err := strconv.Atoi(c.QueryParam("order"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid status order",
		}
	}

	err = h.statusUseCase.DeleteStatus(c.Request().Context(), statusID, order)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't delete status: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusNoContent)
}
