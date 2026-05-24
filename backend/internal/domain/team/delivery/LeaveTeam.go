package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// LeaveTeam handles a user leaving a team.
// @Summary Remove user from team
// @ID			team-leave
// @Tags		team
// @Param		id		path	int	true	"Team id to remove user from"
// @Success	204
// @Failure	400	{object}	echo.HTTPError	"Bad request"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/team/{id}/leave [delete].
func (h *TeamHandler) LeaveTeam(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid team ID")
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Validation("Failed to authenticate user")
	}

	err = h.teamUseCase.RemoveUserFromTeam(c.Request().Context(), teamID, claims.ID)
	if err != nil {
		return apierror.Validation("Failed to remove user")
	}

	return c.NoContent(http.StatusNoContent)
}
