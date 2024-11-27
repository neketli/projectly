package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TODO: fix response type
// @Summary		Get users boards list
// @ID				board-user-get-list
// @Tags			board
// @Accept			json
// @Produce		json
// @Param			user_id	query		int	false	"User id"
// @Success		200		{array} 	interface{}
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/board/list-user [get]
func (ph *BoardHandler) GetUserBoardList(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
		}
	}

	boards, err := ph.boardUseCase.GetUserBoards(c.Request().Context(), userID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get boards: %s", err.Error()),
		}
	}
	return c.JSON(http.StatusOK, boards)
}
