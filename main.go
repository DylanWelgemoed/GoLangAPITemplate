package main

import (
	"log"
	"net/http"
	"os"
)

func main() {	
	port := os.Getenv("PORT")
	
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(port, router))
}