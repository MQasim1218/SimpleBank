package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	db_driver   = "postgres"
	db_location = "postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(db_driver, db_location)
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
