package main

import (
	"github.com/gorilla/handlers"
	"jd/core"
	"jd/models"
	"jd/routes"
	"net/http"
	"os"
)

type T struct {
	A int
}

func main() {
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, routes.Router))
}
