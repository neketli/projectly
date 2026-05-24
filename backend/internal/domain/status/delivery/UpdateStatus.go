package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/status/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

type updateStatusRequest struct {
	Title    string `json:"title"`
	Order    int    `json:"order"`
	OldOrder *int   `json:"old_order"`
	HexColor string `json:"hex_color"`
	BoardID  int    `json:"board_id"`
}

// UpdateStatus handles updating an existing status.
// @Summary Update an existing status
// @ID			status-update
// @Tags		status
// @Accept		application/json
// @Produce		application/json
// @Param		request	body	updateStatusRequest	true	"Status details to update"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/status/{id} [patch].
func (h *StatusHandler) UpdateStatus(c echo.Context) error {
	statusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var request updateStatusRequest
	if bindErr := c.Bind(&request); bindErr != nil {
		return apierror.Validation("Invalid request body")
	}

	err = h.statusUseCase.UpdateStatus(c.Request().Context(), &entity.Status{
		ID:       statusID,
		Title:    request.Title,
		Order:    request.Order,
		HexColor: request.HexColor,
		BoardID:  request.BoardID,
	}, request.OldOrder)
	if err != nil {
		return apierror.Internal("Failed to update status")
	}

	return c.NoContent(http.StatusOK)
}
