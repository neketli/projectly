package middlewares

import (
	teamUseCase "projectly-server/internal/domain/team/usecase"
)

// BoardMiddleware provides middleware for board handlers.
type BoardMiddleware struct {
	teamUseCase teamUseCase.TeamUseCase
}

// New creates a new BoardMiddleware instance.
func New(tu teamUseCase.TeamUseCase) *BoardMiddleware {
	return &BoardMiddleware{teamUseCase: tu}
}
