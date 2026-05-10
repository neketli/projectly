package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is the main application configuration.
type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
	PG   `yaml:"pg"`
	S3   `yaml:"s3"`
	Log  `yaml:"logger"`
	Auth `yaml:"auth"`
}

// App represents application metadata.
type App struct {
	Name    string `env-required:"true" yaml:"name"`
	Version string `env-required:"true" yaml:"version"`
	Mode    string `env:"APP_MODE" env-default:"production"`
}

// HTTP represents HTTP server configuration.
type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

// Log represents logger configuration.
type Log struct {
	Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
}

// PG represents PostgreSQL configuration.
type PG struct {
	PoolMax int    `env-required:"true" yaml:"pool_max"`
	DSN     string `env-required:"true" env:"PG_DSN"`
}

// S3 represents S3 storage configuration.
type S3 struct {
	Secure    bool   `env-default:"false" yaml:"secure"`
	Bucket    string `env-required:"true" yaml:"bucket"`
	Host      string `env-required:"true" env:"S3_HOST"`
	AccessKey string `env-required:"true" env:"S3_ACCESS"`
	SecretKey string `env-required:"true" env:"S3_SECRET"`
}

// Auth represents authentication configuration.
type Auth struct {
	AccessTTL          int    `env-required:"true" yaml:"access_ttl"`
	RefreshTTL         int    `env-required:"true" yaml:"refresh_ttl"`
	AccessSecret       string `env-required:"true" env:"AUTH_ACCESS_SECRET"`
	RefreshSecret      string `env-required:"true" env:"AUTH_REFRESH_SECRET"`
	GoogleAuthProvider struct {
		ClientID     string `env-default:"" env:"AUTH_GOOGLE_CLIENT_ID"`
		ClientSecret string `env-default:"" env:"AUTH_GOOGLE_CLIENT_SECRET"`
		CallbackURL  string `env-default:"" env:"AUTH_GOOGLE_CALLBACK_URL"`
	} `env-required:"false"`
	YandexAuthProvider struct {
		ClientID     string `env-default:"" env:"AUTH_YANDEX_CLIENT_ID"`
		ClientSecret string `env-default:"" env:"AUTH_YANDEX_CLIENT_SECRET"`
		CallbackURL  string `env-default:"" env:"AUTH_YANDEX_CALLBACK_URL"`
	} `env-required:"false"`
}

// New creates a new Config instance.
func New() *Config {
	cfg := &Config{}

	var path string
	flag.StringVar(&path, "config", "./config/config.yml", "Path to config")
	flag.Parse()

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		panic(fmt.Errorf("config error: %w", err))
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		panic(fmt.Errorf("config error: %w", err))
	}

	return cfg
}