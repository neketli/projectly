package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// DeleteProject handles the deletion of a project.
// @Summary Delete project
// @ID			project-delete
// @Tags		project
// @Accept		application/json
// @Produce		application/json
// @Param		id	path	int	true	"Project ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/project/{id} [delete].
func (ph *ProjectHandler) DeleteProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	err = ph.projectUseCase.DeleteProject(c.Request().Context(), projectID)
	if err != nil {
		return apierror.Internal("Failed to delete project")
	}

	return c.NoContent(http.StatusOK)
}
