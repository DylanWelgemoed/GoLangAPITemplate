package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}