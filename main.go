package main

import (
	"log"
	"net/http"
	"os"	
	"github.com/codegangsta/negroni" 
	"github.com/DylanWelgemoed/GoLangAPITemplate/routers"
)

func main() {	
	port := os.Getenv("PORT")
	
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Setup for Router, Endpoints, Handlers and middleware
	router := routers.InitRoutes()
	middleWare := negroni.Classic()
	middleWare.UseHandler(router)

	// Serves API - Creates a new thread and if fails it will log the error.
	log.Fatal(http.ListenAndServe(":" + port, middleWare))
}