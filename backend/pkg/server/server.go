package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"task-tracker-server/pkg/logger"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo   *echo.Echo
	logger logger.Interface
}

func New(engine *echo.Echo, logger logger.Interface) *Server {
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())
	engine.Use(middleware.CORS())

	engine.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	return &Server{
		echo:   engine,
		logger: logger,
	}
}

func (s *Server) Start(port string) error {
	go func() {
		if err := s.echo.Start(net.JoinHostPort("", port)); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.echo.Shutdown(ctx)
}
