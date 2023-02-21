package db

import (
	"database/sql"
	"go_api_tuto/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Could not load system config: ,", err)
		return
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("Could not open database: ,", err)
		return
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
