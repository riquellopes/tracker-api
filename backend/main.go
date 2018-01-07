package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	. "github.com/riquellopes/tracker-api/models"
)

var tracker = Tracker{}

// ValidateTrackerHandler -
func ValidateTrackerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	valid, err := tracker.Exists(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if valid {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// CreateTrackerHandler -
func CreateTrackerHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tracker Tracker

	statusCode := http.StatusOK
	response, _ := json.Marshal(map[string]string{"result": "success"})

	if err := json.NewDecoder(r.Body).Decode(&tracker); err != nil {
		statusCode = http.StatusBadRequest
		response, _ = json.Marshal(map[string]string{"result": "Invalid request payload"})

		log.Println(err.Error())
	}

	if err := tracker.Add(); err != nil {
		statusCode = http.StatusInternalServerError
		response, _ = json.Marshal(map[string]string{"result": err.Error()})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// GetAllTrackerHandler -
func GetAllTrackerHandler(w http.ResponseWriter, r *http.Request) {
	trackers, err := tracker.All()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(trackers)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

// RegisterTrackerHandler -
func RegisterTrackerHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var tracker Tracker

	statusCode := http.StatusOK
	response, _ := json.Marshal(map[string]string{"result": "success"})

	if err := json.NewDecoder(r.Body).Decode(&tracker); err != nil {
		statusCode = http.StatusBadRequest
		response, _ = json.Marshal(map[string]string{"result": "Invalid request payload"})

		log.Println(err.Error())
	}

	if err := tracker.Register(params["id"]); err != nil {
		statusCode = http.StatusInternalServerError
		response, _ = json.Marshal(map[string]string{"result": err.Error()})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func init() {
	var (
		USERNAME = os.Getenv("USERNAME")
		PASSWORD = os.Getenv("PASSWORD")
		SERVER   = os.Getenv("SERVER")
		DATABASE = os.Getenv("DATABASE")
	)

	log.Printf("MongoDB configuration: %s %s %s %s", USERNAME, PASSWORD, SERVER, DATABASE)

	Connection(USERNAME, PASSWORD, SERVER, DATABASE)

	log.Println("MongoDB connected.")
}

func main() {

	route := mux.NewRouter()

	// Api path
	route.HandleFunc("/tracker", GetAllTrackerHandler).Methods("GET")
	route.HandleFunc("/tracker", CreateTrackerHandler).Methods("POST")
	route.HandleFunc("/tracker/{id}", ValidateTrackerHandler).Methods("GET")
	route.HandleFunc("/tracker/{id}", RegisterTrackerHandler).Methods("PUT")

	handler := cors.Default().Handler(route)

	if err := http.ListenAndServe(":5000", handlers.LoggingHandler(os.Stdout, handler)); err != nil {
		log.Fatal(err)
	}
}
