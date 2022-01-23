package http_server

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.server.ListenAndServe()
}
