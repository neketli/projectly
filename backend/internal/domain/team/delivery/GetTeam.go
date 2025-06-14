package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/team/entity"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Get team
// @ID			team-get
// @Tags		team
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		int	true	"Team id"
// @Success	200	{object}	entity.Team
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/team/{id} [get]
func (th *TeamHandler) GetTeam(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil || teamID <= 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}
	isUserInTeam, err := th.teamUseCase.CheckUserInTeam(c.Request().Context(), teamID, claims.ID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't check user in team: %s", err.Error()),
		}
	}
	if !isUserInTeam {
		return &echo.HTTPError{
			Code:    http.StatusForbidden,
			Message: "user is not in team",
		}
	}

	var team entity.Team
	team, err = th.teamUseCase.GetTeam(c.Request().Context(), teamID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get users: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, team)
}
