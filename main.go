package main

func main() {
	server := NewServer(":3000")

	adminMiddlewares := []Middleware{CheckAuth(), Loading()} // Middlewares for admin profile

	server.Handler("/", HandleRoot)
	server.Handler("/api", server.AddMiddlewares(HandleRoot, adminMiddlewares...))
	server.Listen()
}
