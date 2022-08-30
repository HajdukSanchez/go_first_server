package main

import (
	"fmt"
	"net/http"
)

// Function to handle Root endpoint request
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
