package server

import (
	"context"
	"github.com/usmonzodasomon/shortener/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      handler,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
