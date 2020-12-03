package myserver

import (
	"github.com/haobogu/lsframework/log"
	"github.com/haobogu/lsframework/protocol"
	"github.com/haobogu/lsframework/server"
)

// MyServer is an example server
type MyServer struct {
	server.BaseLanguageServer
}

// Completion implements LanguageServer.Completion()
func (m *MyServer) Completion(params protocol.CompletionParams) (protocol.CompletionList, error) {
	log.Info("Trigger my server's completion")
	return protocol.CompletionList{}, nil
}
