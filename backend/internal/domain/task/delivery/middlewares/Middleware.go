package middlewares

import (
	teamUseCase "projectly-server/internal/domain/team/usecase"
)

// TaskMiddleware provides middleware for task handlers.
type TaskMiddleware struct {
	teamUseCase teamUseCase.TeamUseCase
}

// New creates a new TaskMiddleware instance.
func New(tu teamUseCase.TeamUseCase) *TaskMiddleware {
	return &TaskMiddleware{teamUseCase: tu}
}
