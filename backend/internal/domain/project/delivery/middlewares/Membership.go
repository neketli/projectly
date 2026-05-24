package middlewares

import (
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TeamMembership validates that the user belongs to the project's team.
func (m *ProjectMiddleware) TeamMembership() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := token.GetUserClaims(c)
			if err != nil {
				return apierror.Validation("Failed to authenticate user")
			}

			var teamID int

			if projectIDStr := c.Param("id"); projectIDStr != "" {
				projectID, errAtoi := strconv.Atoi(projectIDStr)
				if errAtoi != nil || projectID <= 0 {
					return apierror.Validation("Invalid id")
				}

				teamID, err = m.teamUseCase.GetProjectTeamID(c.Request().Context(), projectID)
				if err != nil {
					return apierror.Internal("Failed to get project team")
				}

				c.Set("project_id", projectID)
			} else if teamIDStr := c.QueryParam("team_id"); teamIDStr != "" {
				teamID, err = strconv.Atoi(teamIDStr)
				if err != nil || teamID <= 0 {
					return apierror.Validation("Invalid team_id")
				}
			} else {
				teams, err := m.teamUseCase.GetTeamByUser(c.Request().Context(), claims.ID)
				if err != nil {
					return apierror.Internal("Failed to get user teams")
				}
				if len(teams) == 0 {
					return apierror.Forbidden("User is not in any team")
				}
				teamID = teams[0].ID
			}

			if teamID > 0 {
				isUserInTeam, err := m.teamUseCase.CheckUserInTeam(c.Request().Context(), teamID, claims.ID)
				if err != nil {
					return apierror.Internal("Failed to verify team membership")
				}
				if !isUserInTeam {
					return apierror.Forbidden("User is not in team")
				}

				c.Set("team_id", teamID)
			}

			c.Set("user_id", claims.ID)

			return next(c)
		}
	}
}
