package handlers

import (
	"encoding/json"
	"net/http"
)

// respond writes the HTTP response with the given response object.
func respond(w http.ResponseWriter, response interface{}) {
	// Set the content type to JSON.
	w.Header().Set("Content-Type", "application/json")

	// Encode the response object and write it to the response writer.
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)

		return
	}
}
