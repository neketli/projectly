package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/project/entity"

	"github.com/labstack/echo/v4"
)

type createProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
	TeamID      int    `json:"team_id"`
}

// @Summary	Create a new project
// @ID			project-create
// @Tags		project
// @Accept		application/json
// @Produce	application/json
// @Param		request	body		createProjectRequest	true	"New project details"
// @Success	201		{object}	entity.Project			"Created project"
// @Failure	400		{object}	echo.HTTPError			"Bad request"
// @Failure	500		{object}	echo.HTTPError			"Internal server error"
// @Router		/project/create [post]
func (ph *ProjectHandler) CreateProject(c echo.Context) error {
	var request createProjectRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	project := &entity.Project{
		ID:          0,
		Title:       request.Title,
		Description: request.Description,
		Code:        request.Code,
		TeamID:      request.TeamID,
	}

	err := ph.projectUseCase.CreateProject(c.Request().Context(), project)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't create project: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusCreated, project)
}
