package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/board/entity"

	"github.com/labstack/echo/v4"
)

type createBoardRequest struct {
	Title     string `json:"title"`
	ProjectID int    `json:"project_id"`
}

// CreateBoard handles the creation of a new board.
// @Summary Create a new board
// @ID			board-create
// @Tags		board
// @Accept		application/json
// @Produce	application/json
// @Param		request	body		createBoardRequest	true	"New board details"
// @Success	201		{object}	entity.Board		"Created board"
// @Failure	400		{object}	echo.HTTPError		"Bad request"
// @Failure	500		{object}	echo.HTTPError		"Internal server error"
// @Router		/board/create [post].
func (h *BoardHandler) CreateBoard(c echo.Context) error {
	var request createBoardRequest
	if err := c.Bind(&request); err != nil {
		return apierror.Validation("Invalid request body")
	}

	board := &entity.Board{
		ID:        0,
		Title:     request.Title,
		ProjectID: request.ProjectID,
	}

	err := h.boardUseCase.CreateBoard(c.Request().Context(), board)
	if err != nil {
		return apierror.Validation("Failed to create board")
	}

	return c.JSON(http.StatusCreated, board)
}
