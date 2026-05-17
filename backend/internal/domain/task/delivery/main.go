package delivery

import (
	"projectly-server/internal/domain/task/delivery/middlewares"
	taskUseCase "projectly-server/internal/domain/task/usecase"
	"projectly-server/internal/domain/team/entity"
	teamUseCase "projectly-server/internal/domain/team/usecase"

	"github.com/labstack/echo/v4"
)

// TaskHandler handles task-related HTTP requests.
type TaskHandler struct {
	taskUseCase taskUseCase.TaskUseCase
}

// New initializes the task handler with routes.
func New(router *echo.Group, b taskUseCase.TaskUseCase, tu teamUseCase.TeamUseCase) {
	handler := &TaskHandler{taskUseCase: b}
	middleware := middlewares.New(tu)

	task := router.Group("/task")
	task.POST("/create", handler.CreateTask, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleDeveloper))
	task.PUT("/:id", handler.UpdateTask, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleDeveloper))
	task.DELETE("/:id", handler.DeleteTask, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleDeveloper))
	task.GET("/:id", handler.GetTask, middleware.TeamMembership())
	task.PATCH("/:id/change-status", handler.UpdateTaskStatus, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleDeveloper))

	task.GET("/list", handler.GetTaskList, middleware.TeamMembership())
	task.GET("/", handler.GetTasks, middleware.TeamMembership())

	task.GET("/:id/attachments", handler.GetAttachments, middleware.TeamMembership())
	task.POST("/:id/create-attachments", handler.CreateAttachment, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleDeveloper))
	task.DELETE("/delete-attachment", handler.DeleteAttachment, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleDeveloper))

	task.GET("/:id/comments", handler.GetComments, middleware.TeamMembership())
	task.POST("/:id/create-comment", handler.CreateComment, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleDeveloper))
	task.DELETE("/:id/delete-comment", handler.DeleteComment, middleware.TeamMembership(), middleware.RequireTeamRole(*entity.RoleDeveloper))

}
