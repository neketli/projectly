package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		PG   `yaml:"pg"`
		Log  `yaml:"logger"`
		Auth `yaml:"auth"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"`
		Version string `env-required:"true" yaml:"version"`
		Mode    string `env:"APP_MODE" env-default:"production"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max"`
		DSN     string `env-required:"true" env:"PG_DSN"`
	}

	Auth struct {
		AccessTTL     int    `env-required:"true" yaml:"access_ttl"`
		RefreshTTL    int    `env-required:"true" yaml:"refresh_ttl"`
		AccessSecret  string `env-required:"true" env:"AUTH_ACCESS_SECRET"`
		RefreshSecret string `env-required:"true" env:"AUTH_REFRESH_SECRET"`
	}
)

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
