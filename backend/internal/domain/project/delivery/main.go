package delivery

import (
	"projectly-server/internal/domain/project/delivery/middlewares"
	projUseCase "projectly-server/internal/domain/project/usecase"
	"projectly-server/internal/domain/team/entity"
	teamUseCase "projectly-server/internal/domain/team/usecase"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectUseCase projUseCase.ProjectUseCase
}

func New(router *echo.Group, pu projUseCase.ProjectUseCase, tu teamUseCase.TeamUseCase) {
	handler := &ProjectHandler{projectUseCase: pu}
	middleware := middlewares.New(tu)

	project := router.Group("/project")
	{
		project.POST("/create", handler.CreateProject, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleEditor))
		project.PATCH("/:id", handler.UpdateProject, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleEditor))
		project.DELETE("/:id", handler.DeleteProject, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleOwner))
		project.GET("/:id", handler.GetProject, middleware.TeamMembership())

		project.GET("/list", handler.GetProjectList, middleware.TeamMembership())
		project.GET("/", handler.GetProjectByCode, middleware.TeamMembership())
	}
}
