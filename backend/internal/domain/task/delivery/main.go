package delivery

import (
	"task-tracker-server/internal/domain/task/usecase"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func New(router *echo.Group, b usecase.TaskUseCase) {
	handler := &TaskHandler{taskUseCase: b}

	task := router.Group("/task")
	{
		task.POST("/create", handler.CreateTask)
		task.PUT("/:id", handler.UpdateTask)
		task.DELETE("/:id", handler.DeleteTask)
		task.GET("/:id", handler.GetTask)
		task.PATCH("/:id/change-status", handler.UpdateTaskStatus)

		task.GET("/list", handler.GetTaskList)
		task.GET("/list-user", handler.GetUserTaskList)
	}
}
