package routers

import (
    "github.com/DylanWelgemoed/GoLangAPITemplate/controllers"
    "github.com/DylanWelgemoed/GoLangAPITemplate/core/authentication"
    negroni "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
    router.HandleFunc(
        "/token-auth", 
        controllers.Login,
    ).Methods("POST")
    
    router.Handle(
        "/refresh-token-auth",
        negroni.New(
            negroni.HandlerFunc(controllers.RefreshToken),
        )).Methods("GET")
    
    router.Handle(
        "/logout",
        negroni.New(
            negroni.HandlerFunc(
                authentication.RequireTokenAuthentication,
            ),
            negroni.HandlerFunc(controllers.Logout),
        )).Methods("GET")
    return router
}