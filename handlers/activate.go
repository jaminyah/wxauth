package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wxauth/datatype"
	"wxauth/platform/dbmgr"
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

	fmt.Printf("Server - email: %s, code: %s\n", form.Email, form.Code)

	isValidCode := redismgr.ValidateCode(form.Email, form.Code)
	fmt.Printf("\nAuth code validated: %v", isValidCode)

	body := map[string]interface{}{"code": 400, "msg": "failed", "email": form.Email}

	var passEncoded string
	if isValidCode == true {
		body = map[string]interface{}{"code": 200, "msg": "ok", "email": form.Email}
		passEncoded = redismgr.FetchPass(form.Email)
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)

	fmt.Printf("\nEncrypted password: %s", passEncoded)

	// Write user object to the database
	passRSA := redismgr.FetchPass(form.Email)
	userDm := datatype.UserDataModel{Email: form.Email, PassRSA: passRSA, UserRole: "Client", Services: "Notify"}
	dbInstance := dbmgr.GetInstance()
	dbInstance.InsertUser(userDm)

}
