package main

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/haobogu/lsframework/log"
	"github.com/haobogu/lsframework/protocol"
	"github.com/sourcegraph/jsonrpc2"
)

func Test_Main(t *testing.T) {
	go main()
	time.Sleep(100 * time.Millisecond)
	conn, err := net.Dial("tcp", "127.0.0.1:34172")
	if err != nil {
		t.Fatal("Dial:", err)
	}

	testClient(context.Background(), t, jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}))
}

func testClient(ctx context.Context, t *testing.T, stream jsonrpc2.ObjectStream) {
	ch := clientHandler{}
	cc := jsonrpc2.NewConn(ctx, stream, &ch)
	defer func() {
		if err := cc.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	var got protocol.InitializeResult
	initParam := protocol.InitializeParams{
		WorkspaceFoldersInitializeParams: protocol.WorkspaceFoldersInitializeParams{
			WorkspaceFolders: []protocol.WorkspaceFolder{{
				URI:  "/Uri",
				Name: "TestWorkspace",
			}},
		},
	}
	if err := cc.Call(ctx, "initialize", initParam, &got); err != nil {
		t.Fatal(err)
	}

	if err := cc.Call(ctx, "completion", initParam, &got); err != nil {
		t.Fatal(err)
	}
	log.Infof("Initialize result: %+v", got)
}

// clientHandler is the "client" handler.
type clientHandler struct {
}

func (h *clientHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
}
