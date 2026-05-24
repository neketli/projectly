package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/board/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

type updateBoardRequest struct {
	Title string `json:"title"`
}

// UpdateBoard handles updating an existing board.
// @Summary Update an existing board
// @ID			board-update
// @Tags		board
// @Accept		application/json
// @Produce	application/json
// @Param		request	body	updateBoardRequest	true	"Board details to update"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/board/{id} [patch].
func (h *BoardHandler) UpdateBoard(c echo.Context) error {
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var request updateBoardRequest
	if bindErr := c.Bind(&request); bindErr != nil {
		return apierror.Validation("Invalid request body")
	}

	err = h.boardUseCase.UpdateBoard(c.Request().Context(), &entity.Board{
		ID:    boardID,
		Title: request.Title,
	})
	if err != nil {
		return apierror.Internal("Failed to update board")
	}

	return c.NoContent(http.StatusOK)
}
