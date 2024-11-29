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
// @Param		id	path	int	true	"Status ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/status/{id} [delete]
func (h *StatusHandler) DeleteStatus(c echo.Context) error {
	statusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	err = h.statusUseCase.DeleteStatus(c.Request().Context(), statusID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't delete status: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
