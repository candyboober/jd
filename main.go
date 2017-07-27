package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"jd/api"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", api.GetVacancyList)
	router.HandleFunc("/{id}/", api.GetVacancy)

	log.Fatal(http.ListenAndServe(":8000", router))
}
