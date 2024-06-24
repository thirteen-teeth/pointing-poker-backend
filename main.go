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
	r.HandleFunc("/join", handlers.JoinSessionHandler).Methods("POST")
	r.HandleFunc("/leave", handlers.LeaveSessionHandler).Methods("POST")
	r.HandleFunc("/vote", handlers.VoteHandler).Methods("POST")
	r.HandleFunc("/results", handlers.ResultsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
