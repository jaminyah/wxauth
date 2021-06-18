package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wxauth/captcha"
	"wxauth/e2ee"
	"wxauth/handlers"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbSource = "./auth.db"
)

func main() {

	os.Remove(dbSource)

	/*
		conn, err := sql.Open(dbDriver, dbSource)
		if err != nil {
			log.Fatal("cannot create database: ", err)
		}
		defer conn.Close()

		dbHandle, err = database.CreateTable(conn)
		if err != nil {
			log.Fatal("cannot create user table.")
		}
	*/

	//router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static").HTTPBox()))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/api/getcaptcha", captcha.GenCaptcha)
	http.HandleFunc("/api/verifycaptcha", captcha.VerifyCaptcha)
	http.HandleFunc("/api/register", handlers.RegisterUser)
	http.HandleFunc("/api/activate", handlers.ActivateUser)
	http.HandleFunc("/api/public", e2ee.SendPublicKey)

	fmt.Println("Server is at :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
