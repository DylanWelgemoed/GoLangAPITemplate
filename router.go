package main

import (
	"net/http"
	"./structs"

    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range structs.routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
}