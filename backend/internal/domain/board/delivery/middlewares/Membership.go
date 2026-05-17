package middlewares

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TeamMembership validates that the user belongs to the board's team.
func (m *BoardMiddleware) TeamMembership() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := token.GetUserClaims(c)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
				}
			}

			var teamID int

			if boardIDStr := c.Param("id"); boardIDStr != "" {
				boardID, errAtoi := strconv.Atoi(boardIDStr)
				if errAtoi != nil || boardID <= 0 {
					return &echo.HTTPError{
						Code:    http.StatusBadRequest,
						Message: "invalid id",
					}
				}

				teamID, err = m.teamUseCase.GetBoardTeamID(c.Request().Context(), boardID)
				if err != nil {
					return &echo.HTTPError{
						Code:    http.StatusInternalServerError,
						Message: fmt.Sprintf("can't get board team id: %s", err.Error()),
					}
				}

				c.Set("board_id", boardID)
			} else if teamIDStr := c.QueryParam("team_id"); teamIDStr != "" {
				teamID, err = strconv.Atoi(teamIDStr)
				if err != nil || teamID <= 0 {
					return &echo.HTTPError{
						Code:    http.StatusBadRequest,
						Message: "invalid team_id",
					}
				}
			} else {
				teams, err := m.teamUseCase.GetTeamByUser(c.Request().Context(), claims.ID)
				if err != nil {
					return &echo.HTTPError{
						Code:    http.StatusInternalServerError,
						Message: fmt.Sprintf("can't get user teams: %s", err.Error()),
					}
				}
				if len(teams) == 0 {
					return &echo.HTTPError{
						Code:    http.StatusForbidden,
						Message: "user is not in any team",
					}
				}
				teamID = teams[0].ID
			}

			if teamID > 0 {
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

				c.Set("team_id", teamID)
			}

			c.Set("user_id", claims.ID)

			return next(c)
		}
	}
}
