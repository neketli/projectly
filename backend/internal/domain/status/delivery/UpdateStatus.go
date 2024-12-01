package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"task-tracker-server/internal/domain/status/entity"

	"github.com/labstack/echo/v4"
)

type updateStatusRequest struct {
	Title    string `json:"title"`
	Order    int    `json:"order"`
	HexColor string `json:"hex_color"`
}

// @Summary	Update an existing status
// @ID			status-update
// @Tags		status
// @Accept		application/json
// @Produce		application/json
// @Param		request	body	updateStatusRequest	true	"Status details to update"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/status/{id} [patch]
func (h *StatusHandler) UpdateStatus(c echo.Context) error {
	statusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var request updateStatusRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	err = h.statusUseCase.UpdateStatus(c.Request().Context(), &entity.Status{
		ID:       statusID,
		Title:    request.Title,
		Order:    request.Order,
		HexColor: request.HexColor,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't update status: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
