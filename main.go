package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.Handle("/", http.FileServer(http.Dir("./views/")))

	http.ListenAndServe(":5000", handlers.LoggingHandler(os.Stdout, route))
}
