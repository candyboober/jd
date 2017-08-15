package handlers

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

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range apiRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(authMiddleware(route.HandlerFunc))

	}
	for _, route := range authRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var authRoutes = []Route{
	Route{
		"GetTokenHandler",
		"GET",
		"/get-token",
		GetTokenHandler,
	},
}

var apiRoutes = []Route{
	Route{
		"GetVacancyList",
		"GET",
		"/",
		ListVacancy,
	},
	Route{
		"GetVacancy",
		"GET",
		"/{id:[0-9]+}",
		RetrieveVacancy,
	},
	Route{
		"CreateVacancy",
		"POST",
		"/",
		CreateVacancy,
	},
	Route{
		"UpdateVacancy",
		"PATCH",
		"/{id:[0-9]+}",
		UpdateVacancy,
	},
	Route{
		"DestroyVacancy",
		"DELETE",
		"/{id:[0-9]+}",
		DestroyVacancy,
	},
}
