package server

import (
	"context"
	"encoding/json"
	"net"
	"strconv"

	"github.com/haobogu/lsframework/log"
	"github.com/haobogu/lsframework/protocol"
	"github.com/pkg/errors"
	"github.com/sourcegraph/jsonrpc2"
)

type LanguageServer interface {
	Completion()
}

type LanguageServerHost struct {
	ServerManager
	LanguageServer
}

// ServerManager is an empty server
type ServerManager struct {
	conn        *Connection
	wd          string
	config      Config
	initialized bool
	port        int
}

// NewBaseServer returns an empty language server
func NewBaseServer(port int, wd string, config Config, processor LanguageServer) *LanguageServerHost {
	p := ServerManager{
		port:        port,
		wd:          wd,
		config:      config,
		initialized: false,
	}
	return &LanguageServerHost{
		ServerManager:  p,
		LanguageServer: processor,
	}
}

// Start starts the server and listen
func (s *LanguageServerHost) Start() error {
	log.Info("Starting server...")
	ctx := context.Background()
	lis, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(s.port)) // any available address
	if err != nil {
		return errors.Wrap(err, "Cannot listen to tcp")
	}
	defer func() {
		if lis == nil {
			return // already closed
		}
		if err = lis.Close(); err != nil {
			log.Fatal("An error occurred when closing tcp connection: ", err)
		}
	}()

	// Listen
	for {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		jsonrpc2.NewConn(ctx, jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}), s)
	}
}

// Handle dilivers incoming requests and notifications
func (s *LanguageServerHost) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	log.Info("handling: ", req.Method)
	// Check whether the server is initialized
	if req.Method != "initialize" && !s.initialized {
		log.Error("the server needs to be initialized")
	}

	// Handle requests and notifications
	switch req.Method {
	case "initialize":
		// TODO: tidy
		log.Info("case initialize")
		var param protocol.InitializeParams
		json.Unmarshal(*req.Params, &param)
		log.Infof("param: %+v", param)
		if s.initialized {
			log.Error("the server is already initialized")
		}
		s.initialized = true
		conn.Reply(ctx, req.ID, protocol.InitializeResult{
			ServerInfo: struct {
				Name    string `json:"name"`
				Version string `json:"version,omitempty"`
			}{
				"name",
				"0.0.1",
			},
		})
	case "completion":
		s.LanguageServer.Completion()
		conn.Reply(ctx, req.ID, nil)
	}
}
