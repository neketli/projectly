package middlewares

import (
	teamUseCase "projectly-server/internal/domain/team/usecase"
)

type BoardMiddleware struct {
	teamUseCase teamUseCase.TeamUseCase
}

func New(teamUseCase teamUseCase.TeamUseCase) *BoardMiddleware {
	return &BoardMiddleware{teamUseCase: teamUseCase}
}
