package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// DeleteTeam handles the deletion of a team.
// @Summary Delete team
// @ID			team-delete
// @Tags		team
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		int	true	"Team ID"
// @Success	200	{object}	echo.HTTPError
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/team/{id} [delete].
func (h *TeamHandler) DeleteTeam(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	err = h.teamUseCase.DeleteTeam(c.Request().Context(), teamID)
	if err != nil {
		return apierror.Internal("Failed to delete team")
	}

	return c.NoContent(http.StatusOK)
}
