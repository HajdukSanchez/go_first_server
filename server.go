package main

import "net/http"

type Server struct {
	port string // Port of the server
}

// Function to create a new server
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

// Function to start listen on the server
func (server *Server) Listen() error {
	err := http.ListenAndServe(server.port, nil)
	if err != nil {
		return err
	}
	return nil
}
