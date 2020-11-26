package main

import (
	"os"

	"github.com/haobogu/lsframework/log"
	"github.com/haobogu/lsframework/server"
)

func main() {
	// Create and start server here
	var s server.LanguageServer
	// Create a server config
	config := server.Config{}
	// Replace your server initializer here
	s = server.NewBaseServer(os.Stdin, os.Stdout, config)

	// Start server, waiting for client's connection
	// Once the connection is established, the server will keep listening the requests and notifications
	if err := s.Start(); err != nil {
		log.Error("The server crashed")
	}
	log.Info("The server is shutting down")
	os.Exit(0)
}
