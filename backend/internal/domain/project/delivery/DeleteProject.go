package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary		Delete project
// @ID			project-delete
// @Tags		project
// @Accept		application/json
// @Produce		application/json
// @Param		id	path	int	true	"Project ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/project/{id} [delete]
func (ph *ProjectHandler) DeleteProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	err = ph.projectUseCase.DeleteProject(c.Request().Context(), projectID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't delete project: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
