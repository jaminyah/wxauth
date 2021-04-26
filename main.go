package main

import (
	"database/sql"
	"log"
	"os"
	"wxauth/platform/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	os.Remove("./auth.db")

	conn, err := sql.Open("sqlite3", "./auth.db")
	if err != nil {
		log.Fatal("cannot create database.")
	}
	defer conn.Close()

	userTbl, err := database.CreateTable(conn)
	if err != nil {
		log.Fatal("cannot initialize database.")
	}

}
