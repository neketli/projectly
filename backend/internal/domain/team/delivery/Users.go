package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/team/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Get team users
// @ID			team-users
// @Tags		team
// @Accept		application/json
// @Produce		application/json
// @Param		id	path		int	true	"Team id to fetch users"
// @Success	200	{array}		entity.TeamUser
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/team/{id}/users [get]
func (th *TeamHandler) Users(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var users []entity.TeamUser

	users, err = th.teamUseCase.GetUsers(c.Request().Context(), teamID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get users: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, users)
}
