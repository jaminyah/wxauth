package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"wxauth/platform/database"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

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

	//api for create captcha
	http.HandleFunc("/api/getcaptcha", genCaptcha)

	//api for verify captcha
	http.HandleFunc("/api/verifycaptcha", verifyCaptcha)

	http.HandleFunc("/api/register", registerUser)

	fmt.Println("Server is at :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
