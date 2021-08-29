package server

import (
	"context"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/config"
	"net/http"
)

const (
	shift = 20 // the necessary shift to convert from megabytes to bytes
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.HttpConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			MaxHeaderBytes: cfg.MaxHeaderMegabytes << shift,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
