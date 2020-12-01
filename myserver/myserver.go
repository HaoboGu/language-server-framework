package myserver

import (
	"github.com/haobogu/lsframework/log"
)

type MyProcessor struct {
}

func (m MyProcessor) Completion() {
	log.Info("my completion")
}
