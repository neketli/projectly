package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/team/entity"
	"projectly-server/internal/domain/user/delivery/token"

	"github.com/labstack/echo/v4"
)

// UserTeams handles retrieval of teams for the current user.
// @Summary Get user teams
// @Description	Get teams list which user is a member
// @ID				team-user-teams
// @Tags			team
// @Accept			application/json
// @Produce			application/json
// @Success		200	{array}		entity.Team		"Teams"
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/team/user [get].
func (h *TeamHandler) UserTeams(c echo.Context) error {
	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Validation("Failed to authenticate user")
	}

	var teams []entity.Team

	teams, err = h.teamUseCase.GetTeamByUser(c.Request().Context(), claims.ID)
	if err != nil {
		return apierror.Internal("Failed to get user teams")
	}

	return c.JSON(http.StatusOK, teams)
}
