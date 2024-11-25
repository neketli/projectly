package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	projectEntity "task-tracker-server/internal/domain/project/entity"

	"github.com/labstack/echo/v4"
)

// @Summary		Get projects list
// @Description	Get projects list of team (with team_id) or user (with user_id)
// @ID				project-get-list
// @Tags			project
// @Accept			json
// @Produce		json
// @Param			team_id	query		int	false	"Team id"
// @Param			user_id	query		int	false	"User id"
// @Success		200		{array}		projectEntity.Project
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/project/list [get]
func (ph *ProjectHandler) GetProjectList(c echo.Context) error {
	teamID, err := strconv.Atoi(c.QueryParam("team_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid team id",
		}
	}

	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
		}
	}

	if teamID == 0 && userID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "team_id or user_id should be provided",
		}
	}

	var projects []projectEntity.Project
	if userID != 0 {
		projects, err = ph.projectUseCase.GetUserProjects(c.Request().Context(), userID)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("can't get projects: %s", err.Error()),
			}
		}
	} else if teamID != 0 {
		projects, err = ph.projectUseCase.GetProjectList(c.Request().Context(), teamID)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("can't get projects: %s", err.Error()),
			}
		}
	}
	return c.JSON(http.StatusOK, projects)
}
