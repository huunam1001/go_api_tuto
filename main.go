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

	/// Load cofig
	config, err := util.LoadConfig(".")

	/// Connect mogo db
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(config.MongoDbUri))
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
		Addr:     config.RedisUrl,
		Password: config.RedisPass, // no password set
		DB:       config.RedisDb,   // use default DB
	})

	server := api.NewServer(store, mongoClient, redis)

	errServer := server.Start(config.SeverAddress)

	if errServer != nil {
		log.Fatal("Could not start server: ,", errServer)
	}

}
