package main

import (
	"log"
	"net/http"

	"jd/handlers"
)

func main() {
	route := handlers.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", route))
}
