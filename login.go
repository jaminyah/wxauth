package main

import (
	"fmt"
	"net/http"
)

var Login = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login POST Submit")
	},
)
