package server

import (
	"github.com/haobogu/lsframework/log"
	"github.com/haobogu/lsframework/protocol"
)

// BaseLanguageServer is an empty server
type BaseLanguageServer struct{}

// Initialize does nothing
func (s *BaseLanguageServer) Initialize(params protocol.InitializeParams) (protocol.InitializeResult, error) {
	log.Info("Initializing base server")
	return protocol.InitializeResult{
		ServerInfo: struct {
			Name    string "json:\"name\""
			Version string "json:\"version,omitempty\""
		}{
			Name:    "base server",
			Version: "0.0.1",
		},
		Capabilities: protocol.ServerCapabilities{},
	}, nil
}

// Initialized does nothing
func (s *BaseLanguageServer) Initialized(params protocol.InitializedParams) error {
	log.Info("Base server received initialized notification")
	return nil
}

// Shutdown does nothing
func (s *BaseLanguageServer) Shutdown() {
	log.Info("Shuting down base server")
}

// Completion does nothing
func (s *BaseLanguageServer) Completion(params protocol.CompletionParams) (protocol.CompletionList, error) {
	log.Info("Triggered base server's completion")
	return protocol.CompletionList{}, nil
}
