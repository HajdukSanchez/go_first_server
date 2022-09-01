package main

import (
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
