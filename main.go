package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var passedMessage string

var dad = []string{
	"Did you hear the one about the guy with the broken hearing aid? Neither did he.",
	"What do you call a fly without wings? A walk.",
	"When my wife told me to stop impersonating a flamingo, I had to put my foot down.",
	"What do you call someone with no nose? Nobody knows.",
	"What time did the man go to the dentist? Tooth hurt-y.",
	"Why can’t you hear a pterodactyl go to the bathroom? The p is silent.",
	"How many optometrists does it take to change a light bulb? 1 or 2? 1... or 2?",
	"I was thinking about moving to Moscow but there is no point Russian into things.",
	"Why does Waldo only wear stripes? Because he doesn't want to be spotted.",
	"Do you know where you can get chicken broth in bulk? The stock market.",
}

// Returns Pong
func getPong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("pong")
}

// Random Dad Joke Switcher -- TODO: place http req to canhazdadjoke
func getJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dj := dad[rand.Intn(10)]
	json.NewEncoder(w).Encode(dj)
}

// Message function
func putMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, passedMessage)
}

func main() {

	customPort := flag.Int("port", 5309, "Define custom port")
	customMessage := flag.String("message", "", "Enter custom message")
	flag.Parse()

	// Init router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/ping", getPong).Methods("GET")
	r.HandleFunc("/v1/joke", getJoke).Methods("GET")

	// CustomPort error handling
	if *customPort <= 1024 || *customPort >= 65534 {
		fmt.Println("Invalid port: please enter a value between 1025 and 65533")
		os.Exit(1)
	}

	if *customMessage != "" {
		passedMessage = *customMessage
		r.HandleFunc("/message", putMessage).Methods("GET")
	}

	port := fmt.Sprintf(":%v", *customPort)
	log.Fatal(http.ListenAndServe(port, r))

}
