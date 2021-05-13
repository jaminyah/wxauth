package handlers

import (
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register User")
}
