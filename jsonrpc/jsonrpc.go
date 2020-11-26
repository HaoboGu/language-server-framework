package jsonrpc

import (
	"context"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/haobogu/lsframework/log"
	"github.com/sourcegraph/jsonrpc2"
	websocketjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
)

var c *websocket.Conn
var resp *http.Response
var jsonrpcConn *jsonrpc2.Conn
var ctx context.Context

// NewClient initialize a jsonrpc2 client using default setting
// REMEMBER to call Close() to close the connection!
func NewClient() *jsonrpc2.Conn {
	ctx = context.Background()
	hb := testHandlerB{}
	var err error
	c, resp, err = websocket.DefaultDialer.Dial("ws://localhost:18360", nil)
	if err != nil {
		log.Fatal(err)
	}
	stream := websocketjsonrpc2.NewObjectStream(c)
	jsonrpcConn = jsonrpc2.NewConn(ctx, stream, &hb)
	return jsonrpcConn
}

// Close closes all connections
func Close() {
	defer resp.Body.Close()
	defer c.Close()
	defer jsonrpcConn.Close()
}

// Call is for testing
func Call() string {
	var result string
	if err := jsonrpcConn.Call(ctx, "test", testParameter{"sss"}, &result); err != nil {
		log.Fatal(err)
	}
	log.Infof("Received response: %s", result)
	return result
}

type testParameter struct {
	URI string `json:"uri"`
}

// testHandlerB is the "client" handler.
type testHandlerB struct {
	mu  sync.Mutex
	got []string
}

// Handle receives notifications from conn
func (h *testHandlerB) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	if req.Notif {
		h.mu.Lock()
		defer h.mu.Unlock()
		log.Info("Handling using testHandlerB")
		h.got = append(h.got, string(*req.Params))
		return
	}
	log.Fatalf("testHandlerB got unexpected request %+v", req)
}
