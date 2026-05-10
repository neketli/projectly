package delivery

import (
	"projectly-server/internal/domain/media/usecase"

	"github.com/labstack/echo/v4"
)

// MediaHandler handles media-related HTTP requests.
type MediaHandler struct {
	mediaUseCase usecase.MediaUseCase
}

// New initializes the media handler with routes.
func New(router *echo.Group, uc usecase.MediaUseCase) {
	handler := &MediaHandler{mediaUseCase: uc}

	media := router.Group("/media")
	media.GET("/*", handler.GetMedia)
}
