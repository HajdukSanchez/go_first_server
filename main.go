package main

func main() {
	server := NewServer(":3000")
	server.Handler("/", HandleRoot)
	server.Handler("/api", HandleAPI)
	server.Listen()
}
