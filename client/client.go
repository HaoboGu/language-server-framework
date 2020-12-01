package client

import (
	"context"
	"net"
	"strconv"

	"github.com/haobogu/lsframework/log"
	"github.com/sourcegraph/jsonrpc2"
)

// LanguageClient is a client which can be used for requesting language server
type LanguageClient struct {
	conn *jsonrpc2.Conn
	ctx  context.Context
}

// NewClient returns an instance of LanguageClient
func NewClient(port int) *LanguageClient {
	ctx := context.Background()
	c, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	ch := clientHandler{}
	jsonrpcConn := jsonrpc2.NewConn(ctx, jsonrpc2.NewBufferedStream(c, jsonrpc2.VSCodeObjectCodec{}), &ch)
	return &LanguageClient{
		conn: jsonrpcConn,
		ctx:  ctx,
	}
}

// Call the server
func (c *LanguageClient) Call(method string, param interface{}, resultPointer interface{}) error {
	if err := c.conn.Call(c.ctx, method, param, resultPointer); err != nil {
		log.Fatal(err)
	}
	return nil
}

// Close closes the connection to the server side
func (c *LanguageClient) Close() error {
	if err := c.conn.Close(); err != nil {
		log.Fatal(err)
	}
	log.Info("The client is closed")
	return nil
}

// clientHandler is the "client" handler.
type clientHandler struct {
}

func (h *clientHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
}
