package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// RemoveUser handles removing a user from a team.
// @Summary Remove user from team
// @ID			team-remove-user
// @Tags		team
// @Param		id		path	int	true	"Team id to remove user from"
// @Param		user_id	path	int	true	"User id to kick from team"
// @Success	204
// @Failure	400	{object}	echo.HTTPError	"Bad request"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/team/{id}/remove-user/{user_id} [delete].
func (h *TeamHandler) RemoveUser(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid team ID")
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return apierror.Validation("Invalid user ID")
	}

	err = h.teamUseCase.RemoveUserFromTeam(c.Request().Context(), teamID, userID)
	if err != nil {
		return apierror.Validation("Failed to remove user")
	}

	return c.NoContent(http.StatusNoContent)
}
