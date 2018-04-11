package routers

import (
    "net/http"
    "github.com/DylanWelgemoed/GoLangAPITemplate/controllers"
    "github.com/DylanWelgemoed/GoLangAPITemplate/core/authentication"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
    router.Handle(
        "/users/",
        negroni.New(
            negroni.HandlerFunc(authentication.RequireTokenAuthentication),
            negroni.HandlerFunc(controllers.GetUsers),
		)).Methods("GET")
		
	
	router.Handle(
		"/users/{userId}",
		negroni.New(
            negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.FindUser),
		)).Methods("GET")
    
    return router
}