package main

import (
	"fmt"
	"net/http"
)

var GetCaptcha = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GetCaptcha GET request.")
	},
)
