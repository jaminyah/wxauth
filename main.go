package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"wxauth/handlers"
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

	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/register", handlers.Register).Methods("POST")
	router.HandleFunc("/api/confirm", handlers.Confirm).Methods("Post")
	router.HandleFunc("/api/forgot", handlers.Forgot).Methods("POST")
	router.HandleFunc("/api/reset", handlers.Reset).Methods("POST")
	router.HandleFunc("/api/logout", handlers.Logout).Methods("POST")
	router.HandleFunc("api/getCaptcha", handlers.GetCaptcha).Methods("GET")
	router.HandleFunc("/api/verifyCaptcha", handlers.VerifyCaptcha).Methods("POST")

	fmt.Println("Server on port: 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}
