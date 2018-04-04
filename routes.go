package main

import ( 
	"github.com/DylanWelgemoed/GoLangAPITemplate/structs"
	"github.com/DylanWelgemoed/GoLangAPITemplate/handlers"
)

var RouteList = structs.Routes{
    structs.Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },
    structs.Route{
        "GetUsers",
        "GET",
        "/users/",
        handlers.GetUsers,
    },
    structs.Route{
        "GetUsers",
        "GET",
        "/users/{userId}",
        handlers.FindUser,
    },
}