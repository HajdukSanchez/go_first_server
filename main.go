package main

import "fmt"

func main() {
	port := "3000"
	server := NewServer(":" + port)
	fmt.Println("Server listening on port:", port)

	adminMiddlewares := []Middleware{CheckAuth(), Loading()} // Middlewares for admin profile

	server.Handler("GET", "/", HandleRoot)
	server.Handler("POST", "/api", server.AddMiddlewares(HandleAPI, adminMiddlewares...))

	server.Listen()
}
