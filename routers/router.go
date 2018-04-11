package routers

import (
    "net/http"
    "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
    router := mux.NewRouter()

    //Create Request Logger Handler
    var handler http.Handler
    handler = LogRequest(handler, route.Name)

    //Set Router Routes
    router = SetAuthenticationRoutes(router)
    router = SetUserRoutes(router)

    return router
}
