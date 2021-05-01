package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

	/*userTbl*/
	_, err = database.CreateTable(conn)
	if err != nil {
		log.Fatal("cannot initialize database.")
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	fmt.Println("Server on port 8090")
	http.ListenAndServe(":8090", nil)
}
