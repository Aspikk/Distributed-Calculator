package server

import (
	"net/http"

	"github.com/Aspikk/Distributed-Calculator/internal/handlers"
)

type Server struct {
	Server *http.Server
}

func New() *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/calculate", handlers.AddExpression)

	return &Server{
		Server: &http.Server{
			Handler: mux,
			Addr:    ":8080",
		},
	}
}
