package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wxauth/datatype"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register User")

	//parse request json body
	decoder := json.NewDecoder(r.Body)
	//var param datatype.ConfigJsonBody
	var form datatype.RegisterForm

	err := decoder.Decode(&form)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	//verify the captcha
	body := map[string]interface{}{"code": 400, "msg": "failed", "email": form.Email}

	// verify password + send email + send cookie token

	fmt.Println(form.Email)
	fmt.Println(form.Password)

	//if Store.Verify(param.Id, param.VerifyValue, true) {
	body = map[string]interface{}{"code": 200, "msg": "ok", "email": form.Email}
	//}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}
