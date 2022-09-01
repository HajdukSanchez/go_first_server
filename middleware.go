package main

import (
	"fmt"
	"net/http"
)

// Middleware to check if data from request is valid
func CheckAuth() Middleware {
	return func(handlerFunction http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fllag := true
			if fllag {
				handlerFunction(w, r) // Call next Middleware
			} else {
				fmt.Fprintf(w, "Authentication failed")
				return // Fail authentication
			}
		}
	}
}
