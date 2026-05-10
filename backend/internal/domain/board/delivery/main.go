package delivery

import (
	"projectly-server/internal/domain/board/delivery/middlewares"
	boardUseCase "projectly-server/internal/domain/board/usecase"
	"projectly-server/internal/domain/team/entity"
	teamUseCase "projectly-server/internal/domain/team/usecase"

	"github.com/labstack/echo/v4"
)

// BoardHandler handles board-related HTTP requests.
type BoardHandler struct {
	boardUseCase boardUseCase.BoardUseCase
}

// New initializes the board handler with routes.
func New(router *echo.Group, b usecase.BoardUseCase) {
	handler := &BoardHandler{boardUseCase: b}
	middleware := middlewares.New(tu)

	board := router.Group("/board")
	board.POST("/create", handler.CreateBoard)
	board.PATCH("/:id", handler.UpdateBoard)
	board.DELETE("/:id", handler.DeleteBoard)
	board.GET("/:id", handler.GetBoard)

	board.GET("/list", handler.GetBoardList)
	board.GET("/list-user", handler.GetUserBoardList)
}
