package delivery

import (
	"projectly-server/internal/domain/team/delivery/middlewares"
	"projectly-server/internal/domain/team/entity"
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
	middleware := middlewares.New(tu)

	team := router.Group("/team")
	{
		team.POST("/create", handler.CreateTeam)
		team.GET("/user", handler.UserTeams)
		team.GET("/roles", handler.GetRoles)

		team.PUT("/:id/update", handler.UpdateTeam, middleware.Membership(), middleware.RequireTeamRole(*entity.RoleOwner))
		team.DELETE("/:id", handler.DeleteTeam, middleware.Membership(), middleware.RequireTeamRole(*entity.RoleOwner))
		team.GET("/:id", handler.GetTeam, middleware.Membership())

		team.GET("/:id/statistic", handler.GetStatistics, middleware.Membership())

		team.GET("/:id/users", handler.Users, middleware.Membership())
		team.POST("/:id/add-user", handler.AddUser, middleware.Membership())
		team.DELETE("/:id/remove-user/:user_id", handler.RemoveUser, middleware.Membership(), middleware.RequireTeamRole(*entity.RoleEditor))
		team.DELETE("/:id/leave", handler.LeaveTeam, middleware.Membership())

		team.POST("/:id/role", handler.SetRole, middleware.Membership(), middleware.RequireTeamRole(*entity.RoleOwner))
	}
}
