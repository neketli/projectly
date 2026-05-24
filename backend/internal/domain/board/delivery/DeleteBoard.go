package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// DeleteBoard handles the deletion of a board.
// @Summary Delete board
// @ID			board-delete
// @Tags		board
// @Accept		application/json
// @Produce	application/json
// @Param		id	path	int	true	"Board ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/board/{id} [delete].
func (h *BoardHandler) DeleteBoard(c echo.Context) error {
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	err = h.boardUseCase.DeleteBoard(c.Request().Context(), boardID)
	if err != nil {
		return apierror.Internal("Failed to delete board")
	}

	return c.NoContent(http.StatusOK)
}
