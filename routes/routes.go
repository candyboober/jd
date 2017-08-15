package routes

import (
	"jd/core"
	"jd/handlers"
	"jd/middlewares"
)

var routes = []core.Route{

	// auth routes
	core.Route{
		"SignIn",
		"POST",
		"/sign-in",
		handlers.SignIn,
	},
	core.Route{
		"SignUp",
		"POST",
		"/sign-up",
		handlers.SignUp,
	},

	// vacancy routes
	core.Route{
		"GetVacancyList",
		"GET",
		"/",
		handlers.ListVacancy,
	},
	core.Route{
		"GetVacancy",
		"GET",
		"/{id:[0-9]+}",
		handlers.RetrieveVacancy,
	},
	core.Route{
		"CreateVacancy",
		"POST",
		"/",
		middlewares.AuthMiddleware(handlers.CreateVacancy),
	},
	core.Route{
		"UpdateVacancy",
		"PATCH",
		"/{id:[0-9]+}",
		handlers.UpdateVacancy,
	},
	core.Route{
		"DestroyVacancy",
		"DELETE",
		"/{id:[0-9]+}",
		handlers.DestroyVacancy,
	},
}

var Router = core.NewRouter(routes)
