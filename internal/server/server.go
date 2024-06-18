package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Aspikk/Distributed-Calculator/internal/handlers"
)

type Server struct {
	Server *http.Server
}

func New() *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/calculate", handlers.AddExpression)
	mux.HandleFunc("/api/v1/expressions", handlers.GetExpressions)
	mux.HandleFunc("/api/v1/expressions/", handlers.GetExpressioinById)

	return &Server{
		Server: &http.Server{
			Handler: mux,
			Addr:    ":8080",
		},
	}
}

func (s *Server) Run() error {
	err := s.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server crushed with err: %v", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.Server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("server shutdowned with error: %w", err)
	}

	return nil
}
