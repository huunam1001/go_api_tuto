package main

import (
	"database/sql"
	"go_api_tuto/api"
	db "go_api_tuto/db/sqlc"
	"go_api_tuto/util"

	_ "github.com/lib/pq"

	"log"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load system config: ,", err)
		return
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("Could not open database: ,", err)
		return
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	errServer := server.Start(config.SeverAddress)

	if errServer != nil {
		log.Fatal("Could not start server: ,", errServer)
	}
}
