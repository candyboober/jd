package main

import (
	"log"
	"net/http"

	"jd/api"
)

func main() {
	router := api.RootRoute
	log.Fatal(http.ListenAndServe(":8000", router))
}
