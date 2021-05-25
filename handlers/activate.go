package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wxauth/datatype"
	"wxauth/redismgr"
)

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Activate")

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

	fmt.Printf("email: %s, code: %s", emailAddr, authCode)

	isValidCode := redismgr.ValidateCode(emailAddr, authCode)
	fmt.Printf("Auth code validated: %v\n", isValidCode)

	body := map[string]interface{}{"code": 400, "msg": "failed", "email": form.Email}
	if isValidCode {
		body = map[string]interface{}{"code": 200, "msg": "ok", "email": form.Email}
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)

	// Save password in db
	if isValidCode {
		// read redis store for password
	}

}
