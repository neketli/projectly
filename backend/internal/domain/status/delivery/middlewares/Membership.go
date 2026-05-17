package middlewares

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TeamMembership validates that the user belongs to the status's team.
func (m *StatusMiddleware) TeamMembership() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var statusID int
			var err error

			if idParam := c.Param("id"); idParam != "" {
				statusID, err = strconv.Atoi(idParam)
				if err != nil || statusID <= 0 {
					return &echo.HTTPError{
						Code:    http.StatusBadRequest,
						Message: "invalid id",
					}
				}
			} else {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "id is required",
				}
			}

			claims, err := token.GetUserClaims(c)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
				}
			}

			teamID, err := m.teamUseCase.GetStatusTeamID(c.Request().Context(), statusID)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("can't get status team id: %s", err.Error()),
				}
			}

			isUserInTeam, err := m.teamUseCase.CheckUserInTeam(c.Request().Context(), teamID, claims.ID)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("can't check user in team: %s", err.Error()),
				}
			}
			if !isUserInTeam {
				return &echo.HTTPError{
					Code:    http.StatusForbidden,
					Message: "user is not in team",
				}
			}

			c.Set("user_id", claims.ID)
			c.Set("team_id", teamID)
			c.Set("status_id", statusID)

			return next(c)
		}
	}
}
