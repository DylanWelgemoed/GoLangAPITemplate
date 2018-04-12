package routers

import (
    //"net/http"
    "log"
    "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
    router := mux.NewRouter()


    //Create Request Logger Handler
    //var handler http.Handler
    //handler = LogRequest(handler, route.Name)

    //Set Router Routes
    router = SetAuthenticationRoutes(router)
    router = SetUserRoutes(router)

    log.Printf("Router intialized...")

    return router
}
