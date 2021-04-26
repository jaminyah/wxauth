package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriver = "sqlite3"
	dbSource = "./auth_test.db"
)

var handle *DbHandle

func TestMain(m *testing.M) {

	os.Remove(dbSource)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot create database: ", err)
	}
	defer conn.Close()

	handle, err = CreateTable(conn)
	if err != nil {
		log.Fatal("cannot create user table.")
	}

	os.Exit(m.Run())

}
