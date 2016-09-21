package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route


func NewRouter(routes Routes) *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    
    //Looping between routes from main.go

    fmt.Println("Adding routes:\n");
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)
        fmt.Println("Adding route: "+route.Name+" -> "+route.Method+" "+route.Pattern);
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    //Return router created
    return router
}
