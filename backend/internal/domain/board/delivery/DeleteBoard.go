package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete board
// @ID			board-delete
// @Tags		board
// @Accept		application/json
// @Produce	application/json
// @Param		id	path	int	true	"Board ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/board/{id} [delete]
func (h *BoardHandler) DeleteBoard(c echo.Context) error {
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	err = h.boardUseCase.DeleteBoard(c.Request().Context(), boardID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't delete board: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
