package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	boardEntity "task-tracker-server/internal/domain/board/entity"

	"github.com/labstack/echo/v4"
)

// @Summary		Get boards list
// @Description	Get boards list of project
// @ID				board-get-list
// @Tags			board
// @Accept			json
// @Produce		json
// @Param			project_id	query		int	false	"Project id"
// @Success		200		{array}		boardEntity.Board
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/board/list [get]
func (ph *BoardHandler) GetBoardList(c echo.Context) error {
	projectID, err := strconv.Atoi(c.QueryParam("project_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid project id",
		}
	}

	var boards []boardEntity.Board
	boards, err = ph.boardUseCase.GetBoardList(c.Request().Context(), projectID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get boards: %s", err.Error()),
		}
	}
	return c.JSON(http.StatusOK, boards)
}
