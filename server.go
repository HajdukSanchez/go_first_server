package main

import "net/http"

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

// Function to start listen on the server
func (server *Server) Listen() error {
	http.Handle("/", server.router) // Handle HTTP request of this endpoint
	err := http.ListenAndServe(server.port, nil)
	if err != nil {
		return err
	}
	return nil
}
