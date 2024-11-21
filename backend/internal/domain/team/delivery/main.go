package delivery

import (
	"task-tracker-server/internal/domain/team/usecase"
	userUseCase "task-tracker-server/internal/domain/user/usecase"

	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
	teamUseCase usecase.TeamUseCase
	userUseCase userUseCase.UserUseCase
}

func New(router *echo.Group, tu usecase.TeamUseCase, u userUseCase.UserUseCase) {
	handler := &TeamHandler{teamUseCase: tu, userUseCase: u}

	team := router.Group("/team")
	{
		team.POST("/create", handler.CreateTeam)
		team.PUT("/update", handler.UpdateTeam)
		team.DELETE("/delete", handler.DeleteTeam)

		team.GET("/user", handler.UserTeams)
		team.GET("/:id/users", handler.Users)
		team.POST("/:id/add-user", handler.AddUser)
		team.DELETE("/:id/remove-user", handler.RemoveUser)
	}
}
