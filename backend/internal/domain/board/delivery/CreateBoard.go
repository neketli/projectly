package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/board/entity"

	"github.com/labstack/echo/v4"
)

type createBoardRequest struct {
	Title     string `json:"title"`
	ProjectID int    `json:"project_id"`
}

// @Summary Create a new board
// @ID			board-create
// @Tags		board
// @Accept		application/json
// @Produce		application/json
// @Param		request	body		createBoardRequest	true	"New board details"
// @Success	201		{object}	entity.Board			"Created board"
// @Failure	400		{object}	echo.HTTPError			"Bad request"
// @Failure	500		{object}	echo.HTTPError			"Internal server error"
// @Router		/board/create [post]
func (ph *BoardHandler) CreateBoard(c echo.Context) error {
	var request createBoardRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	board := &entity.Board{
		ID:        0,
		Title:     request.Title,
		ProjectID: request.ProjectID,
	}

	err := ph.boardUseCase.CreateBoard(c.Request().Context(), board)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't create board: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusCreated, board)
}
