package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"task-tracker-server/internal/domain/board/entity"

	"github.com/labstack/echo/v4"
)

type updateBoardRequest struct {
	Title string `json:"title"`
}

// @Summary	Update an existing board
// @ID			board-update
// @Tags		board
// @Accept		json
// @Produce	json
// @Param		request	body	updateBoardRequest	true	"Board details to update"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/board/{id} [patch]
func (ph *BoardHandler) UpdateBoard(c echo.Context) error {
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var request updateBoardRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	err = ph.boardUseCase.UpdateBoard(c.Request().Context(), &entity.Board{
		ID:    boardID,
		Title: request.Title,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't update board: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
