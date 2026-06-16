package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/shivamx64/shipyard-app/internal/handlers"
)

type Server struct {
	httpServer *http.Server
}

func New(port string, handler *handlers.Handler) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/version", handler.VersionHandler)

	httpServer := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on %s", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server")

	return s.httpServer.Shutdown(ctx)
}