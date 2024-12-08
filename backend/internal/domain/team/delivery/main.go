package delivery

import (
	"projectly-server/internal/domain/team/usecase"
	userUseCase "projectly-server/internal/domain/user/usecase"

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
		team.DELETE("/:id", handler.DeleteTeam)
		team.GET("/:id", handler.GetTeam)

		team.GET("/:id/statistic", handler.GetStatistics)

		team.GET("/user", handler.UserTeams)
		team.GET("/:id/users", handler.Users)
		team.POST("/:id/add-user", handler.AddUser)
		team.DELETE("/:id/remove-user/:user_id", handler.RemoveUser)

		team.GET("/roles", handler.GetRoles)
		team.POST("/:id/role", handler.SetRole)
	}
}
