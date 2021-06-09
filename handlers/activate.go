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

	var passEncoded string
	if isValidCode == true {
		body = map[string]interface{}{"code": 200, "msg": "ok", "email": form.Email}
		passEncoded = redismgr.FetchPass(emailAddr)
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)

	fmt.Printf("\nEncrypted password: %s", passEncoded)

	// TODO - Write password to the user database

	/*
		cipherText, _ := base64.StdEncoding.DecodeString(passEncoded)
		fmt.Print("cipherText: ")
		fmt.Print(cipherText)

		privateKey := e2ee.FetchPriKey()
		data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
		if err != nil {
			log.Println(err)
		}
		fmt.Print("Data: ")
		fmt.Print(string(data))
	*/

	fmt.Println("\nActivate.go - end")
}
