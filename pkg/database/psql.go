package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PostgresConfig struct {
	Host         string `env-required:"true"`
	Port         string `env-required:"true"`
	DatabaseName string `yaml:"databaseName" env-required:"true"`
	User         string `env-required:"true"`
	Password     string `env-required:"true"`
	SSLMode      string `yaml:"ssl_mode" env-required:"true"`
}

func NewPostgresConnection(cf PostgresConfig) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cf.Host, cf.Port, cf.User, cf.Password, cf.DatabaseName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("successfully connected!")

	return db
}
