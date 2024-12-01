package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"task-tracker-server/internal/domain/status/entity"

	"github.com/labstack/echo/v4"
)

// @Summary		Get status list
// @Description	Get status list of project board
// @ID				status-list
// @Tags			status
// @Accept			application/json
// @Produce			application/json
// @Param			board_id	query		int	false	"Board id"
// @Success		200			{array}		entity.Status
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/status/list [get]
func (h *StatusHandler) GetStatusList(c echo.Context) error {
	boardID, err := strconv.Atoi(c.QueryParam("board_id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid board id",
		}
	}

	var statusList []entity.Status
	statusList, err = h.statusUseCase.GetStatusList(c.Request().Context(), boardID)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get status list: %s", err.Error()),
		}
	}
	return c.JSON(http.StatusOK, statusList)
}
