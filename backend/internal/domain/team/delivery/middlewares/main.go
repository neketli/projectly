package middlewares

import "projectly-server/internal/domain/team/usecase"

// TeamMiddleware provides team-related middleware functions.
type TeamMiddleware struct {
	teamUseCase usecase.TeamUseCase
}

// New creates a new TeamMiddleware instance.
func New(u usecase.TeamUseCase) *TeamMiddleware {
	return &TeamMiddleware{teamUseCase: u}
}
