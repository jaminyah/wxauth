package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"wxauth/models"
	"wxauth/platform/database"

	"github.com/mojocn/base64Captcha"

	_ "github.com/mattn/go-sqlite3"
)

var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
func getCaptcha(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Generating captcha")

	decoder := json.NewDecoder(r.Body)
	var param models.ConfigJsonBody

	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	var driver base64Captcha.Driver

	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

func verifyCaptcha(w http.ResponseWriter, r *http.Request) {

	//parse request json body
	decoder := json.NewDecoder(r.Body)
	var param models.ConfigJsonBody

	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	//verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if store.Verify(param.Id, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

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
	/*
		router := mux.NewRouter()
		//router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static").HTTPBox()))
		router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

		router.HandleFunc("/api/login", handlers.Login).Methods("POST")
		router.HandleFunc("/api/register", handlers.Register).Methods("POST")
		router.HandleFunc("/api/confirm", handlers.Confirm).Methods("Post")
		router.HandleFunc("/api/forgot", handlers.Forgot).Methods("POST")
		router.HandleFunc("/api/reset", handlers.Reset).Methods("POST")
		router.HandleFunc("/api/logout", handlers.Logout).Methods("POST")
		router.HandleFunc("/api/getcaptcha", GenCaptcha).Methods("POST")
		router.HandleFunc("/api/verifycaptcha", VerifyCaptcha).Methods("POST")
	*/

	//fmt.Println("Server on port: 8090")
	//log.Fatal(http.ListenAndServe(":8090", router))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	//api for create captcha
	http.HandleFunc("/api/getCaptcha", getCaptcha)

	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", verifyCaptcha)

	fmt.Println("Server is at :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
