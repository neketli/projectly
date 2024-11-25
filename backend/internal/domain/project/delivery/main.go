package delivery

import (
	"task-tracker-server/internal/domain/project/usecase"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectUseCase usecase.ProjectUseCase
}

func New(router *echo.Group, pu usecase.ProjectUseCase) {
	handler := &ProjectHandler{projectUseCase: pu}

	project := router.Group("/project")
	{
		project.POST("/create", handler.CreateProject)
		project.PATCH("/:id", handler.UpdateProject)
		project.DELETE("/:id", handler.DeleteProject)
		project.GET("/:id", handler.GetProject)

		project.GET("/list", handler.GetProjectList)
		project.GET("/", handler.GetProjectByCode)
	}
}
