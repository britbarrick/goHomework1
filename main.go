package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Returns Pong
func getPong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("pong")
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/ping", getPong).Methods("GET")
	// r.HandleFunc("/v1/joke", getJoke).Methods("GET")

	log.Fatal(http.ListenAndServe(":5309", r))
}
