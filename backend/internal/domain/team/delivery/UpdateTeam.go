package delivery

import (
	"fmt"
	"net/http"

	"task-tracker-server/internal/domain/team/entity"

	"github.com/labstack/echo/v4"
)

// @Summary     Update an existing team
// @ID          update-team
// @Tags        team
// @Accept      json
// @Produce     json
// @Param       team body entity.Team true "Team details to update"
// @Success     200 {object} entity.Team "Updated team details"
// @Failure     400 {object} echo.HTTPError "Invalid input"
// @Failure     404 {object} echo.HTTPError "Team not found"
// @Failure     500 {object} echo.HTTPError "Internal server error"
// @Router      /team/update [put]
func (th *TeamHandler) UpdateTeam(c echo.Context) error {
	var team entity.Team
	err := c.Bind(&team)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	err = th.teamUseCase.UpdateTeam(c.Request().Context(), &team)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't update team: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, team)
}
