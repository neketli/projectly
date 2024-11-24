package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type setRoleRequest struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

// @Summary     Set user role in team
// @ID          team-set-role
// @Tags        team
// @Accept      json
// @Produce     json
// @Param       request body setRoleRequest true "user id and role id"
// @Param       id   path     int     true        "Team ID"
// @Success     200
// @Failure     400  {object} echo.HTTPError
// @Failure     500  {object} echo.HTTPError
// @Router      /team/:id/role [post]
func (th *TeamHandler) SetRole(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid team id",
		}
	}

	var request setRoleRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	err = th.teamUseCase.SetRole(c.Request().Context(), teamID, request.UserID, request.RoleID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't set role: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
