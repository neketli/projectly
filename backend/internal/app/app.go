package app

import (
	"fmt"
	"task-tracker-server/config"
	"task-tracker-server/internal/domain/user"
	userDelivery "task-tracker-server/internal/domain/user/delivery"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"
	"task-tracker-server/pkg/server"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.DSN, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	l.Info("Postgres connection is ok!")

	e := echo.New()
	api := e.Group("/api/v1")

	userUsecase := user.New(user.Dependency{
		Config:   cfg,
		Logger:   l,
		Postgres: pg,
	})
	userDelivery.New(api, userUsecase)

	protected := api.Group("")
	protected.Use(echojwt.JWT([]byte(cfg.Auth.AccessSecret)))
	protected.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	server := server.New(e, l)
	server.Start(cfg.HTTP.Port)
}
