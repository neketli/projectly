package middlewares

import (
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/team/entity"

	"github.com/labstack/echo/v4"
)

// RequireTeamRole validates that the user has the required role in the team.
func (m *BoardMiddleware) RequireTeamRole(requiredRole entity.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, ok := c.Get("user_id").(int)
			if !ok {
				return apierror.Validation("Invalid user id")
			}

			teamID, ok := c.Get("team_id").(int)
			if !ok {
				return apierror.Validation("Invalid team id")
			}

			userRole, err := m.teamUseCase.GetUserRole(c.Request().Context(), teamID, userID)
			if err != nil {
				return apierror.Validation("Failed to get user role")
			}

			if userRole == nil || userRole.ID > requiredRole.ID {
				return apierror.Forbidden("User doesn't have required role")
			}

			return next(c)
		}
	}
}
