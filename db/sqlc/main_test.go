package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)


var (
	dbsource = "postgresql://root:secret@localhost:5432/newbank?sslmode=disable"
	dbDriver = "postgres"
)
var testQueries *Queries
var testDB *sql.DB

func TestMain (m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbsource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

