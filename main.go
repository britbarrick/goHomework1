package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

// Returns Pong
func getPong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("pong")
}

// Random Dad Joke Switcher
func getJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dj string
	switch rand.Intn(10) {
	case 0:
		dj = "Did you hear the one about the guy with the broken hearing aid? Neither did he."
	case 1:
		dj = "What do you call a fly without wings? A walk."
	case 2:
		dj = "When my wife told me to stop impersonating a flamingo, I had to put my foot down."
	case 3:
		dj = "What do you call someone with no nose? Nobody knows."
	case 4:
		dj = "What time did the man go to the dentist? Tooth hurt-y."
	case 5:
		dj = "Why can’t you hear a pterodactyl go to the bathroom? The p is silent."
	case 6:
		dj = "How many optometrists does it take to change a light bulb? 1 or 2? 1... or 2?"
	case 7:
		dj = "I was thinking about moving to Moscow but there is no point Russian into things."
	case 8:
		dj = "Why does Waldo only wear stripes? Because he doesn't want to be spotted."
	case 9:
		dj = "Do you know where you can get chicken broth in bulk? The stock market."
	}
	json.NewEncoder(w).Encode(dj)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/ping", getPong).Methods("GET")
	r.HandleFunc("/v1/joke", getJoke).Methods("GET")

	log.Fatal(http.ListenAndServe(":5309", r))
}
