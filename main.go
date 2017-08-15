package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	. "jd/handlers"
	"os"
)

func main() {
	route := NewRouter()
	//log.Fatal(http.ListenAndServe(":8000", route))
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, route))
}
