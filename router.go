package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc // Map to handle request by Method
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// This function is going to search inside Router rules which handler function matches with the request user send
// After that, return the function and also a boolean value to indicate is the key exists in the map[string]http.HandlerFunc
// No we can return another boolean value, this value is going to tell us if the endpoint allow this specific method to be called
// Retun Handler function, methos exists and route exists
func (router *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, routeExist := router.rules[path]
	handler, methodExist := router.rules[path][method]
	return handler, methodExist, routeExist
}

// Function to handle write and read requests serviung it as a response
// It is important to implement this function interface, to use router in Listen function on server.go
func (router *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, routeExist := router.FindHandler(request.URL.Path, request.Method) // Search handler for path
	if !routeExist {
		w.WriteHeader(http.StatusNotFound) // Not found 404
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed) // Not method allowed for this request
		return
	}
	handler(w, request)
}
