package delivery

import (
	"projectly-server/internal/domain/status/usecase"

	"github.com/labstack/echo/v4"
)

// StatusHandler handles status-related HTTP requests.
type StatusHandler struct {
	statusUseCase usecase.StatusUseCase
}

// New initializes the status handler with routes.
func New(router *echo.Group, uc usecase.StatusUseCase) {
	handler := &StatusHandler{statusUseCase: uc}

	status := router.Group("/status")
	status.POST("/create", handler.CreateStatus)
	status.PATCH("/:id", handler.UpdateStatus)
	status.DELETE("/delete", handler.DeleteStatus)
	status.GET("/list", handler.GetStatusList)
}
