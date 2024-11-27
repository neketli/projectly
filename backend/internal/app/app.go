package app

import (
	"context"
	"fmt"
	"task-tracker-server/config"
	"task-tracker-server/internal/domain/project"
	projectDelivery "task-tracker-server/internal/domain/project/delivery"
	"task-tracker-server/internal/domain/team"
	teamDelivery "task-tracker-server/internal/domain/team/delivery"
	"task-tracker-server/internal/domain/user"
	userDelivery "task-tracker-server/internal/domain/user/delivery"
	userEntity "task-tracker-server/internal/domain/user/entity"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/minio"
	"task-tracker-server/pkg/postgres"
	"task-tracker-server/pkg/server"

	"github.com/golang-jwt/jwt/v5"
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

	m, err := minio.New(cfg.S3)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - minio.New: %w", err))
	}
	m.Client.BucketExists(context.Background(), m.Bucket)
	l.Info("Minio connection is ok!")

	e := echo.New()
	api := e.Group("/api/v1")

	auth := api.Group("/auth")

	api.Use(echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(userEntity.JWTClaims)
			},

			SigningKey: []byte(cfg.Auth.AccessSecret),
		},
	))

	userUseCase := user.New(user.Dependency{
		Config:   cfg,
		Logger:   l,
		Postgres: pg,
		S3:       m,
	})
	userDelivery.New(auth, api, userUseCase)

	teamUseCase := team.New(team.Dependency{
		Logger:   l,
		Postgres: pg,
	})
	teamDelivery.New(api, teamUseCase, userUseCase)

	projectUseCase := project.New(project.Dependency{
		Logger:   l,
		Postgres: pg,
	})
	projectDelivery.New(api, projectUseCase)

	server := server.New(e, l)
	server.Start(cfg.HTTP.Port)
}
