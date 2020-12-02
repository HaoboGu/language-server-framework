package myserver

import (
	"github.com/haobogu/lsframework/log"
)

// MyServer is an example server
type MyServer struct {
}

// Completion implements LanguageServer.Completion()
func (m *MyServer) Completion() {
	log.Info("my completion")
}
