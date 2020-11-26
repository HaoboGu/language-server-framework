package server

import (
	"io"

	"github.com/haobogu/lsframework/log"
)

// BaseLanguageServer is an empty server
type BaseLanguageServer struct {
	in     io.Reader
	out    io.Writer
	wd     string
	config Config
}

// NewBaseServer returns an empty language server
func NewBaseServer(in io.Reader, out io.Writer, wd string, config Config) *BaseLanguageServer {
	s := &BaseLanguageServer{
		in:     in,
		out:    out,
		wd:     wd,
		config: config,
	}
	return s
}

// Start starts the server and listen
func (s *BaseLanguageServer) Start() error {
	log.Info("Starting server...")
	return nil
}
