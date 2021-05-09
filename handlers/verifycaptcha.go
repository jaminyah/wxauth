package handlers

import (
	"fmt"
	"net/http"
)

func VerifyCaptcha(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Verify captcha POST Submit")
}
