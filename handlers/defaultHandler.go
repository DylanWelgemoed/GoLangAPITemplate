package handlers

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}