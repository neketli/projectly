package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary		Remove user from team
// @ID			team-leave
// @Tags		team
// @Param		id		path	int	true	"Team id to remove user from"
// @Success	204
// @Failure	400	{object}	echo.HTTPError	"Bad request"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/team/{id}/leave [delete]
func (th *TeamHandler) LeaveTeam(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid team id",
		}
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	err = th.teamUseCase.RemoveUserFromTeam(c.Request().Context(), teamID, claims.ID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't remove user: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusNoContent)
}
