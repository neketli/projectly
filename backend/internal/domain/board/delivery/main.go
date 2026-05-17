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
func New(router *echo.Group, b boardUseCase.BoardUseCase, tu teamUseCase.TeamUseCase) {
	handler := &BoardHandler{boardUseCase: b}
	middleware := middlewares.New(tu)

	board := router.Group("/board")
	board.POST("/create", handler.CreateBoard, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleEditor))
	board.PATCH("/:id", handler.UpdateBoard, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleEditor))
	board.DELETE("/:id", handler.DeleteBoard, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleOwner))
	board.GET("/:id", handler.GetBoard, middleware.TeamMembership())

	board.GET("/list", handler.GetBoardList, middleware.TeamMembership())
	board.GET("/list-user", handler.GetUserBoardList, middleware.TeamMembership())
}
