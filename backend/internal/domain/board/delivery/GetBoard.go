package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	boardEntity "projectly-server/internal/domain/board/entity"

	"github.com/labstack/echo/v4"
)

// GetBoard handles retrieval of a board.
// @Summary Get board
// @ID			board-get
// @Tags		board
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		int	true	"Board id"
// @Success	200	{object}	boardEntity.Board
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/board/{id} [get].
func (h *BoardHandler) GetBoard(c echo.Context) error {
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var board boardEntity.Board
	board, err = h.boardUseCase.GetBoard(c.Request().Context(), boardID)
	if err != nil {
		return apierror.Internal("Failed to get board")
	}

	return c.JSON(http.StatusOK, board)
}
