package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(routes []Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}
	return router
}

var routes = []Route{
	Route{
		"GetVacancyList",
		"GET",
		"/",
		GetVacancyList,
	},
	Route{
		"GetVacancy",
		"GET",
		"/{id}",
		GetVacancy,
	},
}

var RootRoute *mux.Router = NewRouter(routes)
