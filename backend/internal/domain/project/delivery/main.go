package delivery

import (
	"projectly-server/internal/domain/project/usecase"

	"github.com/labstack/echo/v4"
)

// ProjectHandler handles project-related HTTP requests.
type ProjectHandler struct {
	projectUseCase usecase.ProjectUseCase
}

// New initializes the project handler with routes.
func New(router *echo.Group, pu usecase.ProjectUseCase) {
	handler := &ProjectHandler{projectUseCase: pu}

	project := router.Group("/project")
	project.POST("/create", handler.CreateProject)
	project.PATCH("/:id", handler.UpdateProject)
	project.DELETE("/:id", handler.DeleteProject)
	project.GET("/:id", handler.GetProject)

	project.GET("/list", handler.GetProjectList)
	project.GET("/", handler.GetProjectByCode)
}
