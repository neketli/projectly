package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/project/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

type updateProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateProject handles updating an existing project.
// @Summary Update an existing project
// @ID			project-update
// @Tags		project
// @Accept		application/json
// @Produce		application/json
// @Param		request	body	updateProjectRequest	true	"Project details to update"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/project/{id} [patch].
func (ph *ProjectHandler) UpdateProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var request updateProjectRequest
	if bindErr := c.Bind(&request); bindErr != nil {
		return apierror.Validation("Invalid request body")
	}

	err = ph.projectUseCase.UpdateProject(c.Request().Context(), &entity.Project{
		ID:          projectID,
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return apierror.Internal("Failed to update project")
	}

	return c.NoContent(http.StatusOK)
}
