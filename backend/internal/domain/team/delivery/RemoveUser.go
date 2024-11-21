package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type removeUserRequest struct {
	UserID int `json:"user_id"`
}

// @Summary     Remove user from team
// @ID          team-remove-user
// @Tags        team
// @Accept      application/json
// @Produce     application/json
// @Param       id   path     int     true  "Team id to remove user from"
// @Param       request body removeUserRequest true "User id to remove"
// @Success     200
// @Failure     400  {object}  echo.HTTPError "Bad request"
// @Failure     500  {object}  echo.HTTPError "Internal server error"
// @Router      /team/{id}/remove-user [delete]
func (th *TeamHandler) RemoveUser(c echo.Context) error {
	var request removeUserRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	err = th.teamUseCase.AddUserToTeam(c.Request().Context(), teamID, request.UserID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't remove user: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
