package main

import (
	"fmt"
	"net/http"
)

var Confirm = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Confirm POST Submit")
	},
)
