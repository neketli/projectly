package app

import (
	"context"
	"fmt"
	"projectly-server/config"
	"projectly-server/internal/domain/board"
	"projectly-server/internal/domain/project"
	"projectly-server/internal/domain/status"
	"projectly-server/internal/domain/task"
	"projectly-server/internal/domain/team"
	"projectly-server/internal/domain/user"
	userEntity "projectly-server/internal/domain/user/entity"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/minio"
	"projectly-server/pkg/postgres"
	"projectly-server/pkg/server"
	"projectly-server/pkg/validator"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
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
	exists, err := m.Client.BucketExists(context.Background(), m.Bucket)
	if !exists || err != nil {
		l.Error(fmt.Errorf("app - Run - m.Client.BucketExists: %w", err))
	} else {
		l.Info("Minio connection is ok!")
	}

	e := echo.New()
	e.Validator = validator.New()

	api := e.Group("/api/v1")
	auth := api.Group("/auth")

	goth.UseProviders(
		google.New(
			cfg.Auth.GoogleAuthProvider.ClientID,
			cfg.Auth.GoogleAuthProvider.ClientSecret,
			cfg.Auth.GoogleAuthProvider.CallbackUrl,
			"email", "profile",
		),
		github.New(
			cfg.Auth.GithubAuthProvider.ClientID,
			cfg.Auth.GithubAuthProvider.ClientSecret,
			cfg.Auth.GithubAuthProvider.CallbackUrl,
			"user", "user:email",
		),
	)

	api.Use(echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(userEntity.JWTClaims)
			},

			SigningKey: []byte(cfg.Auth.AccessSecret),
		},
	))

	userUseCase := user.New(user.Dependency{
		Config:     cfg,
		Logger:     l,
		Postgres:   pg,
		S3:         m,
		Router:     api,
		AuthRouter: auth,
	})

	team.New(team.Dependency{
		Logger:      l,
		Postgres:    pg,
		Router:      api,
		UserUseCase: userUseCase,
	})

	project.New(project.Dependency{
		Logger:   l,
		Postgres: pg,
		Router:   api,
	})

	board.New(board.Dependency{
		Logger:   l,
		Postgres: pg,
		Router:   api,
	})

	status.New(status.Dependency{
		Logger:   l,
		Postgres: pg,
		Router:   api,
	})

	task.New(task.Dependency{
		Logger:   l,
		Postgres: pg,
		S3:       m,
		Router:   api,
	})

	server := server.New(e, l)
	server.Start(cfg.HTTP.Port)
}
