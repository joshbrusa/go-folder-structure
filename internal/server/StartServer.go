package server

import (
	"fmt"
	"go-folder-structure/internal/handlers"
	"net/http"
)

func (server *Server) StartServer() {
	server.Mux.Handle("/", handlers.RootHandler())

	fmt.Println("server listening on port: ", server.Port)
	http.ListenAndServe(":"+server.Port, server.Mux)
}
