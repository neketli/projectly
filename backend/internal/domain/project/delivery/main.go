package delivery

import (
	"projectly-server/internal/domain/project/delivery/middlewares"
	projUseCase "projectly-server/internal/domain/project/usecase"
	"projectly-server/internal/domain/team/entity"
	teamUseCase "projectly-server/internal/domain/team/usecase"

	"github.com/labstack/echo/v4"
)

// ProjectHandler handles project-related HTTP requests.
type ProjectHandler struct {
	projectUseCase projUseCase.ProjectUseCase
}

// New initializes the project handler with routes.
func New(router *echo.Group, pu usecase.ProjectUseCase) {
	handler := &ProjectHandler{projectUseCase: pu}
	middleware := middlewares.New(tu)

	project := router.Group("/project")
	project.POST("/create", handler.CreateProject)
	project.PATCH("/:id", handler.UpdateProject)
	project.DELETE("/:id", handler.DeleteProject)
	project.GET("/:id", handler.GetProject)

	project.GET("/list", handler.GetProjectList)
	project.GET("/", handler.GetProjectByCode)
}
