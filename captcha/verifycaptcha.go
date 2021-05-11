package captcha

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wxauth/datatype"
)

func VerifyCaptcha(w http.ResponseWriter, r *http.Request) {

	fmt.Println("VerifyCaptcha")

	//parse request json body
	decoder := json.NewDecoder(r.Body)
	var param datatype.ConfigJsonBody

	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	//verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if Store.Verify(param.Id, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}
