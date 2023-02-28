package main

import (
	"context"
	"database/sql"
	"fmt"
	"go_api_tuto/api"
	db "go_api_tuto/db/sqlc"
	"go_api_tuto/util"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://huunam1001:ninhhuunam@cluster0.kcqlk9o.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("MongoDB FAIL")
		log.Fatal(err)
	} else {
		fmt.Println("MongoDB connected")
	}
	defer client.Disconnect(ctx)

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
