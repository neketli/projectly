package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/team/entity"

	"github.com/labstack/echo/v4"
)

// @Summary     Get team roles
// @ID          team-roles
// @Tags        team
// @Accept      json
// @Produce     json
// @Success     200  {array}  entity.Role
// @Failure     400  {object} echo.HTTPError
// @Failure     500  {object} echo.HTTPError
// @Router      /team/roles [get]
func (th *TeamHandler) GetRoles(c echo.Context) error {
	var roles []entity.Role
	roles, err := th.teamUseCase.GetRoles(c.Request().Context())
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get roles: %s", err.Error()),
		}
	}

	return c.JSON(http.StatusOK, roles)
}
