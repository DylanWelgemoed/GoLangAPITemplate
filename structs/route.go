package structs

import (
    "net/http"
    "github.com/DylanWelgemoed/GoLangAPITemplate/handlers"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var RouteList = Routes{
    Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "GetUsers",
        "GET",
        "/users/",
        handlers.GetUsers,
    },
}