package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"projectly-server/internal/domain/board/entity"

	"github.com/labstack/echo/v4"
)

// @Summary	Get users boards list
// @ID			board-user-get-list
// @Tags		board
// @Accept		application/json
// @Produce		application/json
// @Param		user_id	query		int	false	"User id"
// @Success	200		{array}		entity.BoardTeam
// @Failure	400		{object}	echo.HTTPError
// @Failure	500		{object}	echo.HTTPError
// @Router		/board/list-user [get]
func (h *BoardHandler) GetUserBoardList(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
		}
	}

	var boards []entity.BoardTeam
	boards, err = h.boardUseCase.GetUserBoards(c.Request().Context(), userID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get boards: %s", err.Error()),
		}
	}
	return c.JSON(http.StatusOK, boards)
}
