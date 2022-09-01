package main

func main() {
	server := NewServer(":3000")
	server.Handler("/", HandleRoot)
	server.Handler("/api", server.AddMiddlewares(HandleRoot, CheckAuth()))
	server.Listen()
}
