package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

// SessionStore stores the session data in memory
var SessionStore = make([]string, 0)

// NewSessionHandler handles the creation of a new scrum pointing poker session
func NewSessionHandler(w http.ResponseWriter, r *http.Request) {
	// Generate a new session ID
	sessionID := generateSessionID()

	// Store the session ID in memory
	SessionStore = append(SessionStore, sessionID)

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

// CurrentSessionsHandler handles the retrieval of current sessions
func CurrentSessionsHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map for JSON response
	resp := map[string][]string{
		"sessions": SessionStore,
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
