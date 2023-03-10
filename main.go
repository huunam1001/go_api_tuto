package main

import (
	"context"
	"database/sql"
	"go_api_tuto/api"
	db "go_api_tuto/db/sqlc"
	"go_api_tuto/util"
	"time"

	"log"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	/// Connect mogo db
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://huunam1001:ninhhuunam@cluster0.kcqlk9o.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
		return
	}
	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(mongoCtx)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer mongoClient.Disconnect(mongoCtx)

	// err = mognoClient.Ping(mongoCtx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	/// Load cofig
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load system config: ,", err)
		return
	}

	/// connect postgresql
	postgresql, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("Could not open database: ,", err)
		return
	}

	store := db.NewStore(postgresql)

	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "passQu@kh0", // no password set
		DB:       0,            // use default DB
	})

	server := api.NewServer(store, mongoClient, redis)

	errServer := server.Start(config.SeverAddress)

	if errServer != nil {
		log.Fatal("Could not start server: ,", errServer)
	}

}
