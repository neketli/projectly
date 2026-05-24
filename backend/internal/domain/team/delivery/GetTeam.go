package delivery

import (
	"net/http"
	"projectly-server/internal/domain/team/entity"
	"projectly-server/internal/domain/user/delivery/token"
	"projectly-server/pkg/apierror"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetTeam handles retrieval of a team.
// @Summary Get team
// @ID			team-get
// @Tags		team
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		int	true	"Team id"
// @Success	200	{object}	entity.Team
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/team/{id} [get].
func (h *TeamHandler) GetTeam(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil || teamID <= 0 {
		return apierror.Validation("Invalid team ID")
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Unauthorized("Failed to authenticate user")
	}

	isUserInTeam, err := h.teamUseCase.CheckUserInTeam(c.Request().Context(), teamID, claims.ID)
	if err != nil {
		return apierror.Internal("Failed to verify team membership")
	}
	if !isUserInTeam {
		return apierror.Forbidden("You are not a member of this team")
	}

	var team entity.Team
	team, err = h.teamUseCase.GetTeam(c.Request().Context(), teamID)
	if err != nil {
		return apierror.NotFound("Team not found")
	}

	return c.JSON(http.StatusOK, team)
}
