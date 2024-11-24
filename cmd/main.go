package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/SerzhLimon/SongsLib/config"
	serv "github.com/SerzhLimon/SongsLib/internal/transport"
	"github.com/SerzhLimon/SongsLib/pkg/postgres"
	"github.com/SerzhLimon/SongsLib/pkg/postgres/migrations"
)

func main() {
	config := config.LoadConfig()
	db, err := postgres.InitPostgresClient(config.Postgres)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = migrations.Up(db)
	if err != nil {
		log.Fatalf("migration error: %v", err)
	}

	server := serv.NewServer(db)
	routes := serv.ApiHandleFunctions{
		Server: *server,
	}

	router := serv.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}

// viper - config
// zap, logrus - logger
