package delivery

import (
	"task-tracker-server/internal/domain/status/usecase"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	statusUseCase usecase.StatusUseCase
}

func New(router *echo.Group, uc usecase.StatusUseCase) {
	handler := &StatusHandler{statusUseCase: uc}

	status := router.Group("/status")
	{
		status.POST("/create", handler.CreateStatus)
		status.PATCH("/:id", handler.UpdateStatus)
		status.DELETE("/delete", handler.DeleteStatus)
		status.GET("/list", handler.GetStatusList)
	}
}
