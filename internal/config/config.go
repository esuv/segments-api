package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
	"segments-api/pkg/database"
	"time"
)

const (
	LocalEnv = "local"
	DevEnv   = "dev"
	ProdEnv  = "prod"
)

type (
	Config struct {
		Http     HttpConfig
		Postgres database.PostgresConfig
	}

	HttpConfig struct {
		Port           int           `yaml:"port" env-required:"true"`
		MaxHeaderBytes int           `yaml:"maxHeaderBytes" env-required:"true"`
		ReadTimeout    time.Duration `yaml:"readTimeout" env-required:"true"`
		WriteTimeout   time.Duration `yaml:"writeTimeout" env-required:"true"`
	}
)

func MustLoad(configDir string) *Config {
	var cfg Config

	err := cleanenv.ReadConfig(filepath.Join(configDir, Env()+".yml"), &cfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return &cfg
}

func Env() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = LocalEnv
	}

	return appEnv
}
