package delivery

import (
	"projectly-server/internal/domain/status/delivery/middlewares"
	statusUseCase "projectly-server/internal/domain/status/usecase"
	"projectly-server/internal/domain/team/entity"
	teamUseCase "projectly-server/internal/domain/team/usecase"

	"github.com/labstack/echo/v4"
)

// StatusHandler handles status-related HTTP requests.
type StatusHandler struct {
	statusUseCase statusUseCase.StatusUseCase
}

// New initializes the status handler with routes.
func New(router *echo.Group, uc statusUseCase.StatusUseCase, tu teamUseCase.TeamUseCase) {
	handler := &StatusHandler{statusUseCase: uc}
	middleware := middlewares.New(tu)
	status := router.Group("/status")
	
	status.POST("/create", handler.CreateStatus, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleEditor))
	status.PATCH("/:id", handler.UpdateStatus, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleEditor))
	status.DELETE("/delete", handler.DeleteStatus, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleOwner))
	status.GET("/list", handler.GetStatusList, middleware.TeamMembership())
}
