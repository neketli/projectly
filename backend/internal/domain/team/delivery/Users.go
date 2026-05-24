package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/team/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Users handles retrieval of team users.
// @Summary Get team users
// @ID			team-users
// @Tags		team
// @Accept		application/json
// @Produce		application/json
// @Param		id	path		int	true	"Team id to fetch users"
// @Success	200	{array}		entity.TeamUser
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/team/{id}/users [get].
func (h *TeamHandler) Users(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var users []entity.TeamUser

	users, err = h.teamUseCase.GetUsers(c.Request().Context(), teamID)
	if err != nil {
		return apierror.Internal("Failed to get users")
	}

	return c.JSON(http.StatusOK, users)
}
