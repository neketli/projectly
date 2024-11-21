package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/team/entity"
	"task-tracker-server/internal/domain/user/delivery/token"

	"github.com/labstack/echo/v4"
)

// @Summary     Get user teams
// @Description Get teams list which user is a member
// @ID          team-user-teams
// @Tags        team
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.Team "Teams"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     500 {object} echo.HTTPError "Internal server error"
// @Router      /team/user [get]
func (h *TeamHandler) UserTeams(c echo.Context) error {
	claims, err := token.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	var teams []entity.Team

	teams, err = h.teamUseCase.GetTeamByUser(c.Request().Context(), claims.ID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get user teams: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, teams)
}
