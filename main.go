package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ValidateHandler -
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemted")
}

// CreateTracker -
func CreateTracker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemted")
}

func main() {

	route := mux.NewRouter()

	route.Handle("/", http.FileServer(http.Dir("./views/")))

	// Api path
	route.HandleFunc("/tracker/{id}", ValidateHandler).Methods("GET")
	route.HandleFunc("/tracker", CreateTracker).Methods("POST")

	if err := http.ListenAndServe(":5000", handlers.LoggingHandler(os.Stdout, route)); err != nil {
		log.Fatal(err)
	}
}
