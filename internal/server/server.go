package server

import "net/http"

type Server struct {
	Server *http.Server
}

func New() *Server {
	mux := http.NewServeMux()

	return &Server{
		Server: &http.Server{
			Handler: mux,
			Addr:    ":8080",
		},
	}
}
