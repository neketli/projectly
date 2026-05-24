package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	projectEntity "projectly-server/internal/domain/project/entity"

	"github.com/labstack/echo/v4"
)

// GetProject handles retrieval of a project.
// @Summary Get project
// @ID			project-get
// @Tags		project
// @Accept		application/json
// @Produce		application/json
// @Param		id	path		int	true	"Project id"
// @Success	200	{object}	projectEntity.Project
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/project/{id} [get].
func (ph *ProjectHandler) GetProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var project projectEntity.Project
	project, err = ph.projectUseCase.GetProject(c.Request().Context(), projectID)
	if err != nil {
		return apierror.Internal("Failed to get project")
	}

	return c.JSON(http.StatusOK, project)
}
