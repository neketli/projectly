package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"task-tracker-server/internal/domain/team/entity"

	"github.com/labstack/echo/v4"
)

type addUserRequest struct {
	UserEmail string `json:"email"`
}

// @Summary     Add user to team
// @ID          team-add-user
// @Tags        team
// @Accept      application/json
// @Produce     application/json
// @Param       id   path     int     true  "Team id to add user"
// @Param       request body addUserRequest true "User email to invite to team"
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
			Message: fmt.Sprintf("team id validation error: %s", err.Error()),
		}
	}

	user, err := th.userUseCase.GetUserByEmail(c.Request().Context(), request.UserEmail)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't get user: %s", err.Error()),
		}
	}

	err = th.teamUseCase.AddUserToTeam(c.Request().Context(), teamID, user.ID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't add user: %s", err.Error()),
		}
	}

	err = th.teamUseCase.SetRole(c.Request().Context(), teamID, user.ID, entity.RoleUser.ID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't set user role: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusCreated)
}
