package server

import (
	"context"
	"io"

	"github.com/haobogu/lsframework/log"
	"github.com/pkg/errors"
	"github.com/sourcegraph/jsonrpc2"
)

// LanguageServer is an empty server
type LanguageServer struct {
	conn        *Connection
	wd          string
	config      Config
	initialized bool
}

// NewBaseServer returns an empty language server
func NewBaseServer(in io.ReadCloser, out io.WriteCloser, wd string, config Config) *LanguageServer {
	s := &LanguageServer{
		conn:        &Connection{in, out},
		wd:          wd,
		config:      config,
		initialized: false,
	}
	return s
}

// Start starts the server and listen
func (s *LanguageServer) Start() error {
	log.Info("Starting server...")
	ctx := context.Background()
	// Put language server itself as the handler of jsonrpc2 connection
	jsonrpc2Conn := jsonrpc2.NewConn(ctx, jsonrpc2.NewBufferedStream(s.conn, jsonrpc2.VSCodeObjectCodec{}), s)
	<-jsonrpc2Conn.DisconnectNotify()
	if err := s.conn.Close(); err != nil {
		return errors.Wrap(err, "the server is closed unexpected")
	}
	return nil
}

// Handle dilivers incoming requests and notifications
func (s *LanguageServer) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	// Check whether the server is initialized
	if req.Method != "initialize" && !s.initialized {
		log.Error("the server needs to be initialized")
	}

	// Handle requests and notifications
	switch req.Method {
	case "initialize":
		if s.initialized {
			log.Error("the server is already initialized")
		}
	}
}
