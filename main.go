package main

import (
	"log"
	"net/http"

	"jd/handlers"
)

func main() {
	router := handlers.RootRoute
	log.Fatal(http.ListenAndServe(":8000", router))
}
