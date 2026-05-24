package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	projectEntity "projectly-server/internal/domain/project/entity"

	"github.com/labstack/echo/v4"
)

// GetProjectByCode handles retrieval of a project by code.
// @Summary Get project
// @ID			project-get-by-code
// @Tags		project
// @Accept		application/json
// @Produce		application/json
// @Param		team_id	query		int		true	"Team id"
// @Param		code	query		string	true	"Project code name"
// @Success	200		{object}	projectEntity.Project
// @Failure	400		{object}	echo.HTTPError
// @Failure	500		{object}	echo.HTTPError
// @Router		/project [get].
func (ph *ProjectHandler) GetProjectByCode(c echo.Context) error {
	teamID, err := strconv.Atoi(c.QueryParam("team_id"))
	if err != nil {
		return apierror.Validation("Invalid team ID")
	}

	code := c.QueryParam("code")
	if code == "" {
		return apierror.Validation("Invalid project code")
	}

	var project projectEntity.Project
	project, err = ph.projectUseCase.GetProjectByCode(c.Request().Context(), teamID, code)
	if err != nil {
		return apierror.Internal("Failed to get project")
	}

	return c.JSON(http.StatusOK, project)
}
