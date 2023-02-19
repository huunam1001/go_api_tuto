package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driver   = "postgres"
	dbSource = "postgresql://root:pass123@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(driver, dbSource)

	if err != nil {
		log.Fatal("Could not open database: ,", err)
		return
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
