package main

import (
	"fmt"
	"net/http"
)

var Forgot = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Forgot POST Submit")
	},
)
