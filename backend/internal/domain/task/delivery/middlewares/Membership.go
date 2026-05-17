package middlewares

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TeamMembership validates that the user belongs to the task's team.
func (m *TaskMiddleware) TeamMembership() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			taskID, err := strconv.Atoi(c.Param("id"))
			if err != nil || taskID <= 0 {
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

			teamID, err := m.teamUseCase.GetTaskTeamID(c.Request().Context(), taskID)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("can't get task team id: %s", err.Error()),
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
			c.Set("task_id", taskID)

			return next(c)
		}
	}
}
