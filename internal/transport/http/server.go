package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labib0x9/sockforces/config"
	"github.com/labib0x9/sockforces/internal/transport/http/handlers/submissions"
	"github.com/labib0x9/sockforces/internal/transport/http/middlewares"
)

type Server struct {
	server     http.Server
	subHandler submissions.Handler
}

func NewServer(subHandler submissions.Handler) *Server {
	return &Server{subHandler: subHandler}
}

func (s *Server) Start(cnf *config.Config) {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	s.subHandler.RegisterRoutes(mux, manager)

	addr := fmt.Sprintf("http://%s:%d", cnf.Addr, cnf.Port)
	s.server = http.Server{
		Addr:    fmt.Sprintf(":%d", cnf.Port),
		Handler: wrappedMux,
	}

	fmt.Printf("Starting %s Server at %s\n", cnf.Service, addr)
	err := s.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Server ListenAndServe():", "error", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
