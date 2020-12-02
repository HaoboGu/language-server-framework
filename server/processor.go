package server

import "github.com/haobogu/lsframework/log"

// EmptyLanguageServer is an empty server
type EmptyLanguageServer struct{}

// Completion does nothings
func (s *EmptyLanguageServer) Completion() {
	log.Info("empty server's completion")
}
