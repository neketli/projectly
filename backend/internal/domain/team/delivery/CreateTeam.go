package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/team/entity"
	"task-tracker-server/internal/domain/user/delivery/token"

	"github.com/labstack/echo/v4"
)

type createTeamRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// @Summary     Create a new team
// @ID          create-team
// @Tags        team
// @Accept      application/json
// @Produce     application/json
// @Param       request body createTeamRequest true "New team details"
// @Success     201  {object}  entity.Team "Created team"
// @Failure     400  {object}  echo.HTTPError "Bad request"
// @Failure     500  {object}  echo.HTTPError "Internal server error"
// @Router      /team/create [post]
func (th *TeamHandler) CreateTeam(c echo.Context) error {
	var request createTeamRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	team := &entity.Team{
		ID:          0,
		Name:        request.Name,
		Description: request.Description,
	}

	err := th.teamUseCase.CreateTeam(c.Request().Context(), team)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't create team: %s", err.Error()),
		}
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	err = th.teamUseCase.AddUserToTeam(c.Request().Context(), team.ID, claims.ID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't add user %s to team: %s", claims.Email, err.Error()),
		}
	}

	return c.JSON(http.StatusCreated, team)
}
