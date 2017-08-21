package routes

import (
	"jd/core"
	"jd/handlers"
	"jd/middlewares"
	"jd/sockets"
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

	core.Route{
		"Messages",
		"GET",
		"/messages",
		sockets.MakeWsHandler(),
	},
}

var Router = core.NewRouter(routes)
