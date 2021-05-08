package main

import (
	"fmt"
	"net/http"
)

var Logout = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Logout POST Submit")
	},
)
