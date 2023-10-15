package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:61926114@127.0.0.1:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to databse: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}