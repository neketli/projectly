package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"projectly-server/pkg/apierror"
	"projectly-server/pkg/logger"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

// Server provides HTTP server functionality.
type Server struct {
	echo   *echo.Echo
	logger logger.Interface
}

// New creates a new Server instance.
func New(engine *echo.Echo, log logger.Interface) *Server {
	engine.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())
	engine.Use(middleware.CORS())
	engine.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))
	engine.HideBanner = true

	engine.HTTPErrorHandler = func(err error, c echo.Context) {
		var (
			code        int
			kind        apierror.Kind
			message     string
			internalErr error
		)

		var apiErr *apierror.Error
		if errors.As(err, &apiErr) {
			code = apiErr.Kind.HTTPStatus()
			kind = apiErr.Kind
			message = apiErr.Message
			internalErr = apiErr.Err
		} else {
			var httpErr *echo.HTTPError
			if errors.As(err, &httpErr) {
				code = httpErr.Code
				kind = apierror.KindFromHTTPStatus(code)
				message = fmt.Sprintf("%v", httpErr.Message)
				internalErr = httpErr.Internal
			} else {
				code = http.StatusInternalServerError
				kind = apierror.KindInternal
				message = "Internal server error"
				internalErr = err
			}
		}

		fields := logger.Fields{
			"error_kind": kind.String(),
			"method":     c.Request().Method,
			"path":       c.Request().URL.Path,
			"status":     code,
		}
		if internalErr != nil {
			fields["internal_error"] = internalErr.Error()
		}

		if code >= 500 {
			log.WithFields(fields).Error(err.Error())
		} else {
			log.WithFields(fields).Info(err.Error())
		}

		if !c.Response().Committed {
			//nolint:errcheck // nothing to do on write error at this point
			c.JSON(code, apierror.NewResponse(kind, message))
		}
	}

	engine.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	return &Server{
		echo:   engine,
		logger: log,
	}
}

// Start starts the HTTP server.
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
