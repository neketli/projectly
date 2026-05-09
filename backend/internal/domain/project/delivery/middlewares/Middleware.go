package middlewares

import (
	teamUseCase "projectly-server/internal/domain/team/usecase"
)

type ProjectMiddleware struct {
	teamUseCase teamUseCase.TeamUseCase
}

func New(teamUseCase teamUseCase.TeamUseCase) *ProjectMiddleware {
	return &ProjectMiddleware{teamUseCase: teamUseCase}
}
