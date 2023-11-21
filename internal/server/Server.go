package server

import "net/http"

type Server struct {
	Mux  *http.ServeMux
	Port string
}
