package migrations

import (
	"database/sql"
	"embed"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	defaultDriver = "postgres"
)

//go:embed *.sql
var embedMigrations embed.FS

var database *sql.DB

func setup(databaseDsn string) {
	var err error
	database, err = sql.Open(defaultDriver, databaseDsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := database.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(defaultDriver); err != nil {
		log.Fatal(err.Error())
	}
}

func Up(databaseDsn string) {
	if database == nil {
		setup(databaseDsn)
	}

	if err := goose.Up(database, "."); err != nil {
		log.Fatal(err.Error())
	}
}

func Down(databaseDsn string) {
	if database == nil {
		setup(databaseDsn)
	}

	if err := goose.Down(database, "."); err != nil {
		log.Fatal(err.Error())
	}
}
