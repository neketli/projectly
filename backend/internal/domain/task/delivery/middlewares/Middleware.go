package middlewares

import (
	teamUseCase "projectly-server/internal/domain/team/usecase"
)

type TaskMiddleware struct {
	teamUseCase teamUseCase.TeamUseCase
}

func New(teamUseCase teamUseCase.TeamUseCase) *TaskMiddleware {
	return &TaskMiddleware{teamUseCase: teamUseCase}
}
