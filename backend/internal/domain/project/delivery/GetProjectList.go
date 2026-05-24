package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	projectEntity "projectly-server/internal/domain/project/entity"

	"github.com/labstack/echo/v4"
)

// GetProjectList handles retrieval of project list.
// @Summary Get projects list
// @Description	Get projects list of team (with team_id) or user (with user_id)
// @ID				project-get-list
// @Tags			project
// @Accept			application/json
// @Produce		application/json
// @Param			team_id	query		int	false	"Team id"
// @Param			user_id	query		int	false	"User id"
// @Success		200		{array}		projectEntity.Project
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/project/list [get].
func (ph *ProjectHandler) GetProjectList(c echo.Context) error {
	param := c.QueryParam("team_id")
	var (
		teamID int
		userID int
		err    error
	)

	if param != "" {
		teamID, err = strconv.Atoi(param)
		if err != nil {
			return apierror.Validation("Invalid team id")
		}
	}

	param = c.QueryParam("user_id")
	if param != "" {
		userID, err = strconv.Atoi(param)
		if err != nil {
			return apierror.Validation("Invalid user id")
		}
	}

	if teamID == 0 && userID == 0 {
		return apierror.Validation("team_id or user_id must be provided")
	}

	var projects []projectEntity.Project
	if userID != 0 {
		projects, err = ph.projectUseCase.GetUserProjects(c.Request().Context(), userID)
		if err != nil {
			return apierror.Internal("Failed to get projects")
		}
	} else if teamID != 0 {
		projects, err = ph.projectUseCase.GetProjectList(c.Request().Context(), teamID)
		if err != nil {
			return apierror.Internal("Failed to get projects")
		}
	}
	return c.JSON(http.StatusOK, projects)
}
