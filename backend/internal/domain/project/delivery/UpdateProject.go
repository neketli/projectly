package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"task-tracker-server/internal/domain/project/entity"

	"github.com/labstack/echo/v4"
)

type updateProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// @Summary	Update an existing project
// @ID			project-update
// @Tags		project
// @Accept		application/json
// @Produce	application/json
// @Param		request	body	updateProjectRequest	true	"Project details to update"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Invalid input"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/project/{id} [patch]
func (ph *ProjectHandler) UpdateProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var request updateProjectRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	err = ph.projectUseCase.UpdateProject(c.Request().Context(), &entity.Project{
		ID:          projectID,
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't update project: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
