package main

import (
	"database/sql"
	"go_api_tuto/api"
	db "go_api_tuto/db/sqlc"

	_ "github.com/lib/pq"

	"log"
)

const (
	driver   = "postgres"
	dbSource = "postgresql://root:pass123@localhost:5432/simple_bank?sslmode=disable"
)

func main() {

	conn, err := sql.Open(driver, dbSource)

	if err != nil {
		log.Fatal("Could not open database: ,", err)
		return
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	errServer := server.Start("0.0.0.0:8000")

	if errServer != nil {
		log.Fatal("Could not start server: ,", errServer)
	}
}
