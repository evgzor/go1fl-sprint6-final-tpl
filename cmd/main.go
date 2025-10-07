package main

import (
	"log"
	"net/http"
	"os"

	"github.com/evgzor/go1fl-sprint6-final-tpl/internal/server"
)

func main() {

	logFile, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	defer logFile.Close()

	logger := log.New(logFile, "app:", log.Ldate|log.Ltime|log.Lshortfile)

	server := server.MakeServer(logger)

	serverError := server.Start()

	if serverError != nil && serverError != http.ErrServerClosed {
		logger.Fatalf("Server failed to start: %v", serverError)
	}

}
