package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"projectly-server/pkg/logger"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type Server struct {
	echo   *echo.Echo
	logger logger.Interface
}

func New(engine *echo.Echo, logger logger.Interface) *Server {
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())
	engine.Use(middleware.CORS())
	engine.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))
	engine.HideBanner = true

	engine.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	return &Server{
		echo:   engine,
		logger: logger,
	}
}

func (s *Server) Start(port string) {
	go func() {
		if err := s.echo.Start(net.JoinHostPort("", port)); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("shutting down the server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.echo.Shutdown(ctx)
	if err != nil {
		s.logger.Fatal("err shutting down the server", err)
	}
}
