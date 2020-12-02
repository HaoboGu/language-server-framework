package myserver

import (
	"github.com/haobogu/lsframework/log"
)

type MyServer struct {
}

func (m *MyServer) Completion() {
	log.Info("my completion")
}
