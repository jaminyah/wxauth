package main

import (
	"fmt"
	"net/http"
)

var VerifyCaptcha = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Verify captcha POST Submit")
	},
)
