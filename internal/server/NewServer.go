package server

import "net/http"

func NewServer() *Server {
	return &Server{
		Mux:  http.NewServeMux(),
		Port: "8000",
	}
}
