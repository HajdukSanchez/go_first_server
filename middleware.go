package main

import (
	"fmt"
	"net/http"
)

// Middleware to check if data from request is valid
func CheckAuth() Middleware {
	return func(handlerFunction http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fllag := false
			fmt.Println("CheckAuth")
			if fllag {
				handlerFunction(w, r) // Call next Middleware
			} else {
				return // Fail authentication
			}
		}
	}
}
