package middlewares

import "projectly-server/internal/domain/team/usecase"

type TeamMiddleware struct {
	teamUseCase usecase.TeamUseCase
}

func New(u usecase.TeamUseCase) *TeamMiddleware {
	return &TeamMiddleware{teamUseCase: u}
}
