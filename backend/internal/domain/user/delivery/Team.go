package delivery

import (
	"fmt"
	"net/http"
	teamEntity "task-tracker-server/internal/domain/team/entity"
	"task-tracker-server/internal/domain/user/delivery/utils"

	"github.com/labstack/echo/v4"
)

// @Summary     Get user teams
// @Description Get teams list which user is a member
// @ID          user-teams
// @Tags        user
// @Accept      json
// @Produce     json
// @Success     200 {array} teamEntity.Team "Teams"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     500 {object} echo.HTTPError "Internal server error"
// @Router      /user/team [get]
func (h *UserHandler) Team(c echo.Context) error {
	claims, err := utils.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	var teams []teamEntity.Team

	teams, err = h.TeamUsecase.GetTeamByUser(c.Request().Context(), claims.ID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get users: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, teams)
}
