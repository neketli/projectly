package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

type setRoleRequest struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

// SetRole handles setting a user role in a team.
// @Summary Set user role in team
// @ID			team-set-role
// @Tags		team
// @Accept		application/json
// @Produce	application/json
// @Param		request	body	setRoleRequest	true	"user id and role id"
// @Param		id		path	int				true	"Team ID"
// @Success	200
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/team/{id}/role [post].
func (h *TeamHandler) SetRole(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid team ID")
	}

	var request setRoleRequest
	if bindErr := c.Bind(&request); bindErr != nil {
		return apierror.Validation("Invalid request body")
	}

	err = h.teamUseCase.SetRole(c.Request().Context(), teamID, request.UserID, request.RoleID)
	if err != nil {
		return apierror.Internal("Failed to set role")
	}

	return c.NoContent(http.StatusOK)
}
