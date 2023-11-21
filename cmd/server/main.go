package main

import "go-folder-structure/internal/server"

func main() {
	server := server.NewServer()
	server.StartServer()
}
