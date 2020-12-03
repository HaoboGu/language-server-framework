package client

import (
	"testing"

	"github.com/haobogu/lsframework/log"
	"github.com/haobogu/lsframework/protocol"
	"github.com/haobogu/lsframework/server"
)

func Test_Client(t *testing.T) {
	port := 34625
	go startServer(port)
	client := NewClient(port)
	initParam := protocol.InitializeParams{
		WorkspaceFoldersInitializeParams: protocol.WorkspaceFoldersInitializeParams{
			WorkspaceFolders: []protocol.WorkspaceFolder{{
				URI:  "/Uri",
				Name: "TestWorkspace",
			}},
		},
	}

	var got protocol.InitializeResult
	if err := client.Call("initialize", initParam, &got); err != nil {
		log.Error("call failed:", err)
	}
	log.Infof("Initialize result: %+v", got)
	client.Close()
}

func startServer(port int) {
	var s *server.LanguageServerHost
	config := server.Config{}
	s = server.NewServerHost(port, ".", config, &server.BaseLanguageServer{})
	if err := s.Start(); err != nil {
		log.Error("The server crashed")
	}
	log.Info("The server is shutting down")
}
