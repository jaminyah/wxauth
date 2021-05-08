package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"wxauth/platform/database"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"

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

	/*
		fileServer := http.FileServer(http.Dir("./static"))
		http.Handle("/", fileServer)
		fmt.Println("Server on port 8090")
		http.ListenAndServe(":8090", nil)
	*/
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static").HTTPBox()))

	router.Handle("/api/login", Login).Methods("POST")
	router.Handle("/api/register", Register).Methods("POST")
	router.Handle("/api/confirm", Confirm).Methods("Post")
	router.Handle("/api/forgot", Forgot).Methods("POST")
	router.Handle("/api/reset", Reset).Methods("POST")
	router.Handle("/api/logout", Logout).Methods("POST")
	router.Handle("/api/getCaptcha", GetCaptcha).Methods("GET")
	router.Handle("/api/verifyCaptcha", VerifyCaptcha).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", router))
}
