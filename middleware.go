package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

// Middleware to know how much time it takes the execution to run
func Loading() Middleware {
	return func(handlerFunction http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}() // Function is going to be executed at the end
			handlerFunction(w, r) // Call next Middleware
		}
	}
}
