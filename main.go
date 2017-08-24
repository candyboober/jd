package main

import (
	"github.com/gorilla/handlers"
	"jd/routes"
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, routes.Router))
}
