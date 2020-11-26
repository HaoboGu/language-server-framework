package server

import (
	"io"

	"github.com/haobogu/lsframework/log"
)

// BaseLanguageServer is an empty server
type BaseLanguageServer struct {
}

// NewBaseServer returns an empty language server
func NewBaseServer(in io.Reader, out io.Writer, config Config) *BaseLanguageServer {
	s := &BaseLanguageServer{}
	err := s.Init()
	if err != nil {
		log.Error("Init failed: ", err)
		return nil
	}
	return s
}

// Init initializes the server
func (bs *BaseLanguageServer) Init() error {
	return nil
}

// Start starts the server and listen
func (bs *BaseLanguageServer) Start() error {
	return nil
}
