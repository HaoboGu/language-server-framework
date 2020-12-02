package myserver

import (
	"testing"

	"github.com/haobogu/lsframework/client"
	"github.com/haobogu/lsframework/log"
	"github.com/haobogu/lsframework/protocol"
	"github.com/haobogu/lsframework/server"
)

func TestMyServer_Completion(t *testing.T) {
	go startServer(46324)
	client := client.NewClient(46324)
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
	if err := client.Call("completion", nil, nil); err != nil {
		log.Error(err)
	}
	log.Info("Closing")
	client.Close()
}

func startServer(port int) {
	s := server.NewBaseServer(port, ".", server.Config{}, &MyServer{})
	if err := s.Start(); err != nil {
		log.Error("The server crashed")
	}
	log.Info("The server is shutting down")
}
