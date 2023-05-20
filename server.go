package TaskManager

import (
	"context"
	"net/http"
	"time"
)

const (
	serverTimeout        = 10 * time.Second
	serverMaxHeaderBytes = 1 << 20
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: serverMaxHeaderBytes,
		ReadTimeout:    serverTimeout,
		WriteTimeout:   serverTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
