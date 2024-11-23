package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/SerzhLimon/SongsLib/config"
	"github.com/SerzhLimon/SongsLib/pkg/postgres"
)

func main() {
	config := config.LoadConfig()
	_, err := postgres.InitPostgresClient(config.Postgres)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("GREAT!")
}
