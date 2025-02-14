package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type ConfigServer struct {
	Port              string
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func (s *Server) RunServer(cfg ConfigServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         cfg.Port,
		Handler:      handler,
		ReadTimeout:  cfg.ReadHeaderTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
