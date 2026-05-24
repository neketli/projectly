package middlewares

import (
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Membership returns a middleware that checks if the user is a member of the team.
func (m *TeamMiddleware) Membership() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			teamID, err := strconv.Atoi(c.Param("id"))
			if err != nil || teamID <= 0 {
				return apierror.Validation("Invalid id")
			}

			claims, err := token.GetUserClaims(c)
			if err != nil {
				return apierror.Validation("Failed to authenticate user")
			}
			isUserInTeam, err := m.teamUseCase.CheckUserInTeam(c.Request().Context(), teamID, claims.ID)
			if err != nil {
				return apierror.Internal("Failed to verify team membership")
			}
			if !isUserInTeam {
				return apierror.Forbidden("User is not in team")
			}

			c.Set("user_id", claims.ID)
			c.Set("team_id", teamID)

			return next(c)
		}
	}
}
