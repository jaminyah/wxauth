package handlers

import (
	"fmt"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout POST Submit")
}
