package handlers

import (
	"fmt"
	"net/http"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reset POST Submit")
}
