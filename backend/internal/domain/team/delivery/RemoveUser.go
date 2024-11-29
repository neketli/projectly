package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Remove user from team
// @ID			team-remove-user
// @Tags		team
// @Accept		application/json
// @Produce	application/json
// @Param		id		path	int	true	"Team id to remove user from"
// @Param		user_id	path	int	true	"User id to kick from team"
// @Success	200
// @Failure	400	{object}	echo.HTTPError	"Bad request"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/team/{id}/remove-user/{user_id} [delete]
func (th *TeamHandler) RemoveUser(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid team id",
		}
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
		}
	}

	err = th.teamUseCase.RemoveUserFromTeam(c.Request().Context(), teamID, userID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't remove user: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
