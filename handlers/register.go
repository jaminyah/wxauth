package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"wxauth/codegen"
	"wxauth/datatype"
	"wxauth/mailmgr"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register User")

	//parse request json body
	decoder := json.NewDecoder(r.Body)

	var form datatype.RegisterForm

	err := decoder.Decode(&form)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	mailAddr := form.Email
	userPass := form.Password
	numDigits := 6

	actcode, err := codegen.GenActCode(numDigits)
	if err != nil {
		log.Println(err)
	}

	response := mailmgr.Send(mailAddr, actcode)
	fmt.Printf("Mail response: %v", response)

	body := map[string]interface{}{"code": response, "msg": "failed", "email": form.Email}

	if response == 202 && isPasswdValid(userPass) {
		body = map[string]interface{}{"code": response, "msg": "ok", "email": form.Email}
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)

	/* User credentials to database */

}

func isPasswdValid(passwd string) bool {

	var isValid bool = true

	regex := regexp.MustCompile(".*[0-9].*")
	regex2 := regexp.MustCompile(".*[a-z].*")
	regex3 := regexp.MustCompile(".*[A-Z].*")
	regex4 := regexp.MustCompile(".*[a-zA-Z0-9]{8,}.*")

	if !regex.MatchString(passwd) {
		isValid = false
	} else if !regex2.MatchString(passwd) {
		isValid = false
	} else if !regex3.MatchString(passwd) {
		isValid = false
	} else if !regex4.MatchString(passwd) {
		isValid = false
	}

	return isValid
}
