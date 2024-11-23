package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/SerzhLimon/SongsLib/config"
	serv "github.com/SerzhLimon/SongsLib/internal/transport"
	"github.com/SerzhLimon/SongsLib/pkg/postgres"
)

func main() {
	config := config.LoadConfig()
	db, err := postgres.InitPostgresClient(config.Postgres)
	if err != nil {
		log.Panic(err)
	}


	server := serv.NewServer(db)
	routes := serv.ApiHandleFunctions{
		Server: *server,
	}

	router := serv.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}
