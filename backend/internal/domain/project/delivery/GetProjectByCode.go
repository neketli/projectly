package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	projectEntity "task-tracker-server/internal/domain/project/entity"

	"github.com/labstack/echo/v4"
)

// @Summary		Get project
// @ID			project-get-by-code
// @Tags		project
// @Accept		application/json
// @Produce		application/json
// @Param		team_id	query		int		true	"Team id"
// @Param		code	query		string	true	"Project code name"
// @Success	200		{object}	projectEntity.Project
// @Failure	400		{object}	echo.HTTPError
// @Failure	500		{object}	echo.HTTPError
// @Router		/project [get]
func (ph *ProjectHandler) GetProjectByCode(c echo.Context) error {
	teamID, err := strconv.Atoi(c.QueryParam("team_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid team id",
		}
	}

	code := c.QueryParam("code")
	if code == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid code",
		}
	}

	var project projectEntity.Project
	project, err = ph.projectUseCase.GetProjectByCode(c.Request().Context(), teamID, code)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get project: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, project)
}
