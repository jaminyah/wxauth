package main

import (
	"fmt"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register POST Submit")
}
