package middlewares

import (
	teamUseCase "projectly-server/internal/domain/team/usecase"
)

// StatusMiddleware provides middleware for status handlers.
type StatusMiddleware struct {
	teamUseCase teamUseCase.TeamUseCase
}

// New creates a new StatusMiddleware instance.
func New(tu teamUseCase.TeamUseCase) *StatusMiddleware {
	return &StatusMiddleware{teamUseCase: tu}
}
