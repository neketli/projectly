// Package app configures and runs application.
package app

import (
	"task-tracker-server/config"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/server"

	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	server := server.New(echo.New(), l)
	server.Start(cfg.HTTP.Port)
}
