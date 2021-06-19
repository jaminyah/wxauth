package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wxauth/datatype"
	"wxauth/e2ee"
	"wxauth/platform/dbmgr"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Server - Login POST Submit")

	//parse request json body
	decoder := json.NewDecoder(r.Body)

	var form datatype.LoginForm

	err := decoder.Decode(&form)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	dbInstance := dbmgr.GetInstance()
	userDataModel, err := dbInstance.GetUser(form.Email)
	if err != nil {
		fmt.Println("Get user error.")
	}
	fmt.Printf("login - db passwrd: %s", userDataModel.PassRSA)
	// fmt.Printf("login - login form.Password: %s", form.Password)

	loginPass := e2ee.DecodeRSA(form.Password)
	dbPass := e2ee.DecodeRSA(userDataModel.PassRSA)

	body := map[string]interface{}{"code": 400, "msg": "failed", "email": form.Email}

	if loginPass == dbPass {
		fmt.Println("Login success.")
		body = map[string]interface{}{"code": 200, "msg": "ok", "email": form.Email}
	}

	// set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)

}
