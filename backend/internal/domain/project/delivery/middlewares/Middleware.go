package middlewares

import (
	teamUseCase "projectly-server/internal/domain/team/usecase"
)

// ProjectMiddleware provides middleware for project handlers.
type ProjectMiddleware struct {
	teamUseCase teamUseCase.TeamUseCase
}

// New creates a new ProjectMiddleware instance.
func New(tu teamUseCase.TeamUseCase) *ProjectMiddleware {
	return &ProjectMiddleware{teamUseCase: tu}
}
