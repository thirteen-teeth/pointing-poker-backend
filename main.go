package main

import (
	"log"
	"net/http"
	"pointing-poker-backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/newsession", handlers.NewSessionHandler).Methods("POST")
	r.HandleFunc("/currentsessions", handlers.CurrentSessionsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
