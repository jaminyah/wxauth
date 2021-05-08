package main

import (
	"fmt"
	"net/http"
)

var Reset = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Reset POST Submit")
	},
)
