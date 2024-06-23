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
	r.HandleFunc("/join/{sessionID}/{userID}", handlers.JoinSessionHandler).Methods("POST")
	r.HandleFunc("/leave/{sessionID}/{userID}", handlers.LeaveSessionHandler).Methods("POST")
	r.HandleFunc("/vote/{sessionID}/{userID}/{vote}", handlers.VoteHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
