package migrations

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed sql/*.sql
var embedMigrations embed.FS

func Up(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		err = fmt.Errorf("%w", err)
		return err
	}

	if err := goose.Up(db, "sql"); err != nil {
		err = fmt.Errorf("%w", err)
		return err
	}

	return nil
}

func Down(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		err = fmt.Errorf("%w", err)
		return err
	}

	if err := goose.Down(db, "sql"); err != nil {
		err = fmt.Errorf("%w", err)
		return err
	}

	return nil
}
