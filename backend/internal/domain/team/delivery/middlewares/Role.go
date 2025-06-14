package middlewares

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/team/entity"

	"github.com/labstack/echo/v4"
)

func (m *TeamMiddleware) RequireTeamRole(requiredRole entity.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, ok := c.Get("user_id").(int)
			if !ok {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid user id",
				}
			}

			teamID, ok := c.Get("team_id").(int)
			if !ok {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid user id",
				}
			}

			userRole, err := m.teamUseCase.GetUserRole(c.Request().Context(), teamID, userID)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("can't get user role: %s", err.Error()),
				}
			}

			if userRole.ID > requiredRole.ID {
				return &echo.HTTPError{
					Code:    http.StatusForbidden,
					Message: "user doesn't have required role",
				}
			}

			return next(c)
		}
	}
}
