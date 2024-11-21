package delivery

import (
	"task-tracker-server/internal/domain/team/usecase"

	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
	teamUseCase usecase.TeamUsecase
}

func New(router *echo.Group, usecase usecase.TeamUsecase) {
	handler := &TeamHandler{teamUseCase: usecase}

	team := router.Group("/team")
	{
		team.POST("/create", handler.CreateTeam)
		team.PUT("/update", handler.UpdateTeam)
		team.DELETE("/delete", handler.DeleteTeam)
		team.GET("/:id/users", handler.Users)
	}
}
