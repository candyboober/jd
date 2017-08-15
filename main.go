package main

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"jd/routes"
)

type T struct {
	A int
}

func main() {
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, routes.Router))
}
