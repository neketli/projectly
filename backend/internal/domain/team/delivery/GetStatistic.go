package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/team/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary		Get team projects statistic
// @Description	Get teams projects statistic
// @ID				team-get-statistic
// @Tags			team
// @Accept			application/json
// @Produce			application/json
// @Param			team_id	query		int	true	"Team id"
// @Success		200		{object}	entity.StatisticData
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/team/{id}/statistic [get]
func (th *TeamHandler) GetStatistics(c echo.Context) error {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
	}

	var statistic []entity.StatisticData
	statistic, err = th.teamUseCase.GetStatisticData(c.Request().Context(), teamID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get statistic data: %s", err.Error()),
		}
	}
	return c.JSON(http.StatusOK, statistic)
}
