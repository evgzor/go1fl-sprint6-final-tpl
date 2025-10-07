package server

import (
	"log"
	"net/http"
	"time"

	"github.com/evgzor/go1fl-sprint6-final-tpl/internal/handlers"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func (s *Server) Start() error {
	s.logger.Println("Start Server on", s.server.Addr)
	return s.server.ListenAndServe()
}

func MakeServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.DownloadHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{logger: logger, server: server}
}
