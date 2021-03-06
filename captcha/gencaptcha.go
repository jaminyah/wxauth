package captcha

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wxauth/datatype"

	"github.com/mojocn/base64Captcha"
)

//var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
func GenCaptcha(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GenCaptcha")

	decoder := json.NewDecoder(r.Body)
	var param datatype.ConfigJsonBody

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

	c := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}
