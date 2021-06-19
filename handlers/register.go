package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"wxauth/codegen"
	"wxauth/datatype"
	"wxauth/e2ee"

	//"wxauth/mailmgr"
	"wxauth/redismgr"
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

	//mailAddr := form.Email
	//passEncoded := form.Password

	numDigits := 6
	actCode, err := codegen.GenActCode(numDigits)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("actcode: %v\n", actCode)
	// fmt.Printf("Register - encoded pass: %s", passEncoded)

	// DEBUG - TEMPORARY SEND NO MAIL
	//response := mailmgr.Send(mailAddr, actcode)
	//fmt.Printf("Mail response: %v", response)

	mailSrvResp := 202 // DEV TESTING
	body := map[string]interface{}{"code": mailSrvResp, "msg": "failed", "email": form.Email}

	redismgr.StoreEmailPass(form.Email, form.Password)
	redismgr.StoreEmailCode(form.Email, actCode)

	passEncoded := redismgr.FetchPass(form.Email)
	userPass := e2ee.DecodeRSA(passEncoded)
	if mailSrvResp == 202 && isPasswdValid(userPass) {
		body = map[string]interface{}{"code": mailSrvResp, "msg": "ok", "email": form.Email}
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
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
