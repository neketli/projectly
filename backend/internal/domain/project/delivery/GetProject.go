package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	projectEntity "projectly-server/internal/domain/project/entity"

	"github.com/labstack/echo/v4"
)

// @Summary		Get project
// @ID			project-get
// @Tags		project
// @Accept		application/json
// @Produce		application/json
// @Param		id	path		int	true	"Project id"
// @Success	200	{object}	projectEntity.Project
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/project/{id} [get]
func (ph *ProjectHandler) GetProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var project projectEntity.Project
	project, err = ph.projectUseCase.GetProject(c.Request().Context(), projectID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get project: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, project)
}
