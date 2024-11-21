package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type addUserRequest struct {
	UserID int `json:"user_id"`
}

// @Summary     Add user to team
// @ID          team-add-user
// @Tags        team
// @Accept      application/json
// @Produce     application/json
// @Param       id   path     int     true  "Team id to add user"
// @Param       request body addUserRequest true "User id to add"
// @Success     201
// @Failure     400  {object}  echo.HTTPError "Bad request"
// @Failure     500  {object}  echo.HTTPError "Internal server error"
// @Router      /team/{id}/add-user [post]
func (th *TeamHandler) AddUser(c echo.Context) error {
	var request addUserRequest
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
			Message: fmt.Sprintf("can't add user: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusCreated)
}
