package middlewares

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (m *ProjectMiddleware) TeamMembership() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			projectID, err := strconv.Atoi(c.Param("id"))
			if err != nil || projectID <= 0 {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid id",
				}
			}

			claims, err := token.GetUserClaims(c)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
				}
			}

			teamID, err := m.teamUseCase.GetProjectTeamID(c.Request().Context(), projectID)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("can't get project team id: %s", err.Error()),
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
			c.Set("project_id", projectID)

			return next(c)
		}
	}
}
