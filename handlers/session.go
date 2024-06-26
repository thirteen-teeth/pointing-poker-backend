package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// SessionStore stores the session data in memory
var SessionStore = make(map[string][]string)

// NewSessionHandler handles the creation of a new scrum pointing poker session
func NewSessionHandler(w http.ResponseWriter, r *http.Request) {
	// Generate a new session ID
	sessionID := generateSessionID()

	// Store the session ID in memory
	SessionStore[sessionID] = []string{}

	// Create a map for JSON response
	resp := map[string]string{
		"message":   "New session created",
		"sessionID": sessionID,
	}

	// Set content-type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Send the response
	json.NewEncoder(w).Encode(resp)
}

// JoinSessionHandler handles a user joining a session
func JoinSessionHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session ID and user ID from the POST data
	sessionID := r.FormValue("sessionID")
	userID := r.FormValue("userID")
	// Add the user to the session
	SessionStore[sessionID] = append(SessionStore[sessionID], userID)
	// Create a map for JSON response
	resp := map[string]string{
		"message":   "User joined session",
		"sessionID": sessionID,
		"userID":    userID,
	}
	// Set content-type to application/json
	w.Header().Set("Content-Type", "application/json")
	// Send the response
	json.NewEncoder(w).Encode(resp)
}

// LeaveSessionHandler handles a user leaving a session
func LeaveSessionHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session ID and user ID from the URL
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]
	userID := vars["userID"]

	// Remove the user from the session
	for i, user := range SessionStore[sessionID] {
		if user == userID {
			SessionStore[sessionID] = append(SessionStore[sessionID][:i], SessionStore[sessionID][i+1:]...)
			break
		}
	}

	// Create a map for JSON response
	resp := map[string]string{
		"message":   "User left session",
		"sessionID": sessionID,
		"userID":    userID,
	}

	// Set content-type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Send the response
	json.NewEncoder(w).Encode(resp)
}

// generateSessionID generates a new session ID
func generateSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// VoteHandler handles a user voting in a session
func VoteHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session ID, user ID, and vote from the URL
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]
	userID := vars["userID"]
	vote := vars["vote"]

	// Create a map for JSON response
	resp := map[string]string{
		"message":   "Vote submitted",
		"sessionID": sessionID,
		"userID":    userID,
		"vote":      vote,
	}

	// Set content-type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Send the response
	json.NewEncoder(w).Encode(resp)
}

// ResultsHandler handles getting the results of a session
func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session ID from the URL
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	// Get the list of users in the session
	users := SessionStore[sessionID]

	// Create a map for JSON response
	resp := map[string]interface{}{
		"sessionID": sessionID,
		"results":   make(map[string]string),
		"users":     users,
	}

	// Set content-type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Send the response
	json.NewEncoder(w).Encode(resp)
}
