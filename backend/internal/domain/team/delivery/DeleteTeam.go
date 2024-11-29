package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete team
// @ID			team-delete
// @Tags		team
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		int	true	"Team ID"
// @Success	200	{object}	echo.HTTPError
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router		/team/{id} [delete]
func (th *TeamHandler) DeleteTeam(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	err = th.teamUseCase.DeleteTeam(c.Request().Context(), teamID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't delete team: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
