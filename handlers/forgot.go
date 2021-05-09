package handlers

import (
	"fmt"
	"net/http"
)

func Forgot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Forgot POST Submit")
}
