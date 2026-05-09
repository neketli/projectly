package middlewares

import (
	teamUseCase "projectly-server/internal/domain/team/usecase"
)

type StatusMiddleware struct {
	teamUseCase teamUseCase.TeamUseCase
}

func New(teamUseCase teamUseCase.TeamUseCase) *StatusMiddleware {
	return &StatusMiddleware{teamUseCase: teamUseCase}
}
