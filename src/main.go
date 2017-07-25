package main

import (
	"log"
	"net/http"
	
	"api"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})

	http.HandleFunc("/", api.GetVacancy)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
