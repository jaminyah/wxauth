package main

import (
	"fmt"
	"net/http"
)

var Register = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Register POST Submit")
	},
)
