package routers

import (
	"net/http"
	"simple-json-restapi/handlers"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes []Route

func init() {

	routes = append(routes, Route{
		Name:        "Index",
		Methods:     []string{"GET"},
		Pattern:     "/",
		HandlerFunc: handlers.Index,
	})

	routes = append(routes, Route{
		Name:        "User",
		Methods:     []string{"DELETE", "GET", "POST"},
		Pattern:     "/user/{id}",
		HandlerFunc: handlers.HandleUser,
	})

	routes = append(routes, Route{
		Name:        "Users",
		Methods:     []string{"GET"},
		Pattern:     "/users",
		HandlerFunc: handlers.HandleUsers,
	})
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
