package main

import (
	"errors"
	"fmt"
	"segments-api/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const configsDir = "configs"
const migrationsDir = "migrations"

func main() {
	cfg := config.MustLoad(configsDir).Postgres

	sourceURL := "file://" + migrationsDir
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Password,
		cfg.DatabaseName,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		cfg.SSLMode,
	)

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Printf("there are no changes yet, add some new migration to ./%s directory\n", migrationsDir)
			return
		}

		panic(err)
	}
}
