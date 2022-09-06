package main

import (
	"net/http"
	"strings"
)

type Server struct {
	port   string // Port of the server
	router *Router
}

// Function to create a new server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Function to connect a path with a handler function
func (server *Server) Handler(method string, path string, handler http.HandlerFunc) {
	_, exist := server.router.rules[path] // Check if this first level key on the map exists
	if !exist {
		server.router.rules[path] = make(map[string]http.HandlerFunc) // Create map
	}
	server.router.rules[path][strings.ToUpper(method)] = handler
}

// This functions allow us to add multiple middlewares to a specific handler function
// Using 3 dots before the DataType, we can pass zero, one or N data of this specific type
func (server *Server) AddMiddlewares(handlerFunction http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handlerFunction = middleware(handlerFunction) // Passing handler function to be executed in the middleware
	}
	return handlerFunction // Here we return the handler function of the path, not a middleware
}

// Function to start listen on the server
func (server *Server) Listen() error {
	http.Handle("/", server.router) // Handle HTTP request of this endpoint
	err := http.ListenAndServe(server.port, nil)
	if err != nil {
		return err
	}
	return nil
}
