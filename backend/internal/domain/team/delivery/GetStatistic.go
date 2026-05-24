package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/team/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetStatistics handles retrieval of team project statistics.
// @Summary Get team projects statistic
// @Description	Get teams projects statistic
// @ID				team-get-statistic
// @Tags			team
// @Accept			application/json
// @Produce			application/json
// @Param			team_id	query		int	true	"Team id"
// @Success		200		{object}	entity.StatisticData
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/team/{id}/statistic [get].
func (h *TeamHandler) GetStatistics(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid ID")
	}

	var statistic []entity.StatisticData
	statistic, err = h.teamUseCase.GetStatisticData(c.Request().Context(), teamID)
	if err != nil {
		return apierror.Internal("Failed to get statistic data")
	}
	return c.JSON(http.StatusOK, statistic)
}
