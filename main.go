package main

import (
	"github.com/gorilla/handlers"
	"jd/routes"
	"net/http"
	"os"
	"log"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, routes.Router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
