package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"wxauth/captcha"
	"wxauth/handlers"
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

	//router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static").HTTPBox()))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/api/getcaptcha", captcha.GenCaptcha)
	http.HandleFunc("/api/verifycaptcha", captcha.VerifyCaptcha)
	http.HandleFunc("/api/register", handlers.RegisterUser)

	fmt.Println("Server is at :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
