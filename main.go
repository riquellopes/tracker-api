package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	. "github.com/riquellopes/tracker-api/models"
)

// ValidateHandler -
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemted")
}

// CreateTracker -
func CreateTracker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemted")
}

func init() {
	var (
		USERNAME = os.Getenv("USERNAME")
		PASSWORD = os.Getenv("PASSWORD")
		SERVER   = os.Getenv("SERVER")
		DATABASE = os.Getenv("DATABASE")
	)

	log.Printf("MongoDB configuration: %s %s", SERVER, DATABASE)

	Connection(USERNAME, PASSWORD, SERVER, DATABASE)

	log.Println("MongoDB connected.")
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
