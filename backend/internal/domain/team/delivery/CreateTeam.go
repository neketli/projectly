package delivery

import (
	"net/http"
	"projectly-server/internal/domain/team/entity"
	"projectly-server/internal/domain/user/delivery/token"
	"projectly-server/pkg/apierror"

	"github.com/labstack/echo/v4"
)

type createTeamRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateTeam handles the creation of a new team.
// @Summary Create a new team
// @ID			create-team
// @Tags		team
// @Accept		application/json
// @Produce	application/json
// @Param		request	body		createTeamRequest	true	"New team details"
// @Success	201		{object}	entity.Team			"Created team"
// @Failure	400		{object}	echo.HTTPError		"Bad request"
// @Failure	500		{object}	echo.HTTPError		"Internal server error"
// @Router		/team/create [post].
func (h *TeamHandler) CreateTeam(c echo.Context) error {
	var request createTeamRequest
	if err := c.Bind(&request); err != nil {
		return apierror.Validation("Invalid request body")
	}

	team := &entity.Team{
		ID:          0,
		Name:        request.Name,
		Description: request.Description,
	}

	err := h.teamUseCase.CreateTeam(c.Request().Context(), team)
	if err != nil {
		return apierror.Internal("Failed to create team")
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Unauthorized("Failed to authenticate user")
	}

	err = h.teamUseCase.AddUserToTeam(c.Request().Context(), team.ID, claims.ID)
	if err != nil {
		return apierror.Internal("Failed to add user to team")
	}

	err = h.teamUseCase.SetRole(c.Request().Context(), team.ID, claims.ID, entity.RoleOwner.ID)
	if err != nil {
		return apierror.Internal("Failed to set owner role")
	}

	return c.JSON(http.StatusCreated, team)
}
