package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wxauth/datatype"
	"wxauth/redismgr"
)

type activateForm struct {
	Email string
	Code  string
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nActivate")

	//parse request json body
	decoder := json.NewDecoder(r.Body)

	var form datatype.ActivateForm

	err := decoder.Decode(&form)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	emailAddr := form.Email
	authCode := form.Code

	fmt.Printf("Server - email: %s, code: %s\n", emailAddr, authCode)

	isValidCode := redismgr.ValidateCode(emailAddr, authCode)
	fmt.Printf("\nAuth code validated: %v", isValidCode)

	body := map[string]interface{}{"code": 400, "msg": "failed", "email": form.Email}
	if isValidCode == true {
		body = map[string]interface{}{"code": 200, "msg": "ok", "email": form.Email}
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)

	// Save password in db
	//if isValidCode {
	// read redis store for password
	//}
	fmt.Println("Activate.go - end")
}
