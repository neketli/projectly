package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	boardEntity "task-tracker-server/internal/domain/board/entity"

	"github.com/labstack/echo/v4"
)

// @Summary	Get board
// @ID			board-get
// @Tags		board
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		int	true	"Board id"
// @Success	200	{object}	boardEntity.Board
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/board/{id} [get]
func (h *BoardHandler) GetBoard(c echo.Context) error {
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var board boardEntity.Board
	board, err = h.boardUseCase.GetBoard(c.Request().Context(), boardID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get board: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, board)
}
