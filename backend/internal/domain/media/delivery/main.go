package delivery

import (
	"projectly-server/internal/domain/media/usecase"

	"github.com/labstack/echo/v4"
)

type MediaHandler struct {
	mediaUseCase usecase.MediaUseCase
}

func New(router *echo.Group, uc usecase.MediaUseCase) {
	handler := &MediaHandler{mediaUseCase: uc}

	media := router.Group("/media")
	{
		media.GET("/*", handler.GetMedia)
	}
}
