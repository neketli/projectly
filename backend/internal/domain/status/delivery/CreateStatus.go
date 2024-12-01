package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/status/entity"

	"github.com/labstack/echo/v4"
)

type createStatusRequest struct {
	BoardID  int    `json:"board_id"`
	Title    string `json:"title"`
	Order    int    `json:"order"`
	HexColor string `json:"hex_color"`
}

// @Summary	Create a new status
// @ID			status-create
// @Tags		status
// @Accept		application/json
// @Produce		application/json
// @Param		request	body		createStatusRequest	true	"New status details"
// @Success	201		{object}	entity.Status		"Created status"
// @Failure	400		{object}	echo.HTTPError		"Bad request"
// @Failure	500		{object}	echo.HTTPError		"Internal server error"
// @Router		/status/create [post]
func (h *StatusHandler) CreateStatus(c echo.Context) error {
	var request createStatusRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	status := &entity.Status{
		ID:       0,
		BoardID:  request.BoardID,
		Title:    request.Title,
		Order:    request.Order,
		HexColor: request.HexColor,
	}

	err := h.statusUseCase.CreateStatus(c.Request().Context(), status)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't create status: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusCreated, status)
}
