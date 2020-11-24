package server

import (
	"io"

	"github.com/haobogu/lsframework/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// BaseLanguageServer is an empty server
type BaseLanguageServer struct {
	logger *zap.Logger
}

// NewBaseServer returns an empty language server
func NewBaseServer(in io.Reader, out io.Writer, config ServerConfig) *BaseLanguageServer {
	logger := log.NewLogger()
	s := &BaseLanguageServer{
		logger: logger
	}
	err := s.Init()
	if err != nil {
		s.logger.Error("Init failed: ", err)
		return errors.Wrap(err)
	}
	return s
}

// Init initializes the server
func (bs *BaseLanguageServer) Init() error {

}

// Start starts the server and listen
func (bs *BaseLanguageServer) Start() error {
	return nil
}
