package delivery

import (
	"errors"
	"net/http"
	"projectly-server/internal/domain/team/entity"
	userEntity "projectly-server/internal/domain/user/entity"
	"projectly-server/pkg/apierror"
	"strconv"

	"github.com/labstack/echo/v4"
)

type addUserRequest struct {
	UserEmail string `json:"email"`
}

// AddUser handles adding a user to a team.
// @Summary Add user to team
// @ID			team-add-user
// @Tags		team
// @Accept		application/json
// @Produce	application/json
// @Param		id		path	int				true	"Team id to add user"
// @Param		request	body	addUserRequest	true	"User email to invite to team"
// @Success	201
// @Failure	400	{object}	echo.HTTPError	"Bad request"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/team/{id}/add-user [post].
func (h *TeamHandler) AddUser(c echo.Context) error {
	var request addUserRequest
	if err := c.Bind(&request); err != nil {
		return apierror.Validation("Invalid request body")
	}

	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid team ID")
	}

	user, err := h.userUseCase.GetUserByEmail(c.Request().Context(), request.UserEmail)
	if err != nil {
		if errors.Is(err, userEntity.ErrNoUserFound) {
			return apierror.NotFound("User not found")
		}
		return apierror.Internal("Failed to find user")
	}

	err = h.teamUseCase.AddUserToTeam(c.Request().Context(), teamID, user.ID)
	if err != nil {
		return apierror.Internal("Failed to add user to team")
	}

	err = h.teamUseCase.SetRole(c.Request().Context(), teamID, user.ID, entity.RoleUser.ID)
	if err != nil {
		return apierror.Internal("Failed to set user role")
	}

	return c.NoContent(http.StatusCreated)
}
