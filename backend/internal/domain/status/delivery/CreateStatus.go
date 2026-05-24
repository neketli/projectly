package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/status/entity"

	"github.com/labstack/echo/v4"
)

type createStatusRequest struct {
	BoardID  int    `json:"board_id"`
	Title    string `json:"title"`
	Order    int    `json:"order"`
	HexColor string `json:"hex_color"`
}

// CreateStatus handles the creation of a new status.
// @Summary Create a new status
// @ID			status-create
// @Tags		status
// @Accept		application/json
// @Produce		application/json
// @Param		request	body		createStatusRequest	true	"New status details"
// @Success	201		{object}	entity.Status		"Created status"
// @Failure	400		{object}	echo.HTTPError		"Bad request"
// @Failure	500		{object}	echo.HTTPError		"Internal server error"
// @Router		/status/create [post].
func (h *StatusHandler) CreateStatus(c echo.Context) error {
	var request createStatusRequest
	if err := c.Bind(&request); err != nil {
		return apierror.Validation("Invalid request body")
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
		return apierror.Validation("Failed to create status")
	}

	return c.JSON(http.StatusCreated, status)
}
