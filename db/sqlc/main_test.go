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
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(db_driver, db_location)
	if err != nil {
		log.Fatal("could not connect to the database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
