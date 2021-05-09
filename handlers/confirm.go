package handlers

import (
	"fmt"
	"net/http"
)

func Confirm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Confirm POST Submit")
}
