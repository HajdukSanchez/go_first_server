package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Function to handle '/' request
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// Function to handle '/api' request
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world from API!")
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body) // get Json from body
	var user User
	// Mapping data from request to a generic interface
	if error := decoder.Decode(&user); error != nil {
		fmt.Fprintf(w, "error %v", error)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") // Send format to header
	w.Write(response)                                  // Response Json format
}
