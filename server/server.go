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

// LanguageServer defines all functions of a server
type LanguageServer interface {
	Initialize(params protocol.InitializeParams) (protocol.InitializeResult, error)
	Initialized(params protocol.InitializedParams) error
	Shutdown()
	Completion(params protocol.CompletionParams) (protocol.CompletionList, error)
}

// LanguageServerHost equals manager + server
type LanguageServerHost struct {
	LanguageServerManager
	LanguageServer
}

// LanguageServerManager contains non-lsp-functional parts, such as connection, configuration, etc.
type LanguageServerManager struct {
	conn        *Connection
	wd          string
	config      Config
	initialized bool
	port        int
}

// NewServerHost returns an empty language server
func NewServerHost(port int, wd string, config Config, server LanguageServer) *LanguageServerHost {
	p := LanguageServerManager{
		port:        port,
		wd:          wd,
		config:      config,
		initialized: false,
	}
	return &LanguageServerHost{
		LanguageServerManager: p,
		LanguageServer:        server,
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
	// Check whether the server is initialized
	if req.Method != "initialize" && !s.initialized {
		log.Error("the server needs to be initialized")
	}

	// Handle requests and notifications
	switch req.Method {
	case "initialize":
		// Check
		if s.initialized {
			log.Error("the server is already initialized")
		}
		if req.Params == nil {
			// TODO: reply error message
			log.Error("Invalid initialize params, ignored")
			return
		}

		// Parse params
		var params protocol.InitializeParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			log.Error("Cannot parse initialize params: ", err)
			return
		}

		// Do initialize, return the initialize result
		initializeResult, err := s.LanguageServer.Initialize(params)
		if err != nil {
			log.Error("Failed to initialize language server: ", err)
			return
		}
		s.initialized = true
		conn.Reply(ctx, req.ID, initializeResult)
	case "initialized":
		var params protocol.InitializedParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			log.Error("Cannot parse initialized params: ", err)
			return
		}
		if err := s.LanguageServer.Initialized(params); err != nil {
			log.Error("Failed to execute initialized for language server")
		}
	case "shutdown":
		s.LanguageServer.Shutdown()
	case "exit":
		if err := conn.Close(); err != nil {
			log.Error("Failed to close jsonrpc2 connection: ", err)
		}
	case "completion":
		if req.Params == nil {
			log.Error("Invalid completion params, ignored")
			return
		}

		// Parse params
		var params protocol.CompletionParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			log.Error("Cannot parse completion params: ", err)
			return
		}

		// Do completion
		result, err := s.LanguageServer.Completion(params)
		if err != nil {
			log.Error("Failed to calculate completion list: ", err)
			return
		}
		conn.Reply(ctx, req.ID, result)
	}
}
