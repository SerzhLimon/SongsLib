package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SerzhLimon/SongsLib/config"
)

func InitPostgresClient(cfg config.PostgresConfig) (*sql.DB, error) {

	options := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.SSLMode)

	database, err := sql.Open("postgres", options)
	if err != nil {
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		return nil, err
	}
	log.Printf("Successful connect to postgres")

	return database, nil
}
