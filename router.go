package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc // Map to handle request
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

// This function is going to search inside Router rules which handler function matches with the request user send
// After that, return the function and also a boolean value to indicate is the key exists in the map[string]http.HandlerFunc
func (router *Router) FindHandler(path string) (http.Handler, bool) {
	handler, exist := router.rules[path]
	return handler, exist
}

// Function to handle write and read requests serviung it as a response
// It is important to implement this function interface, to use router in Listen function on server.go
func (router *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
