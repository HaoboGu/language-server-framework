package server

import "github.com/haobogu/lsframework/log"

type BaseProcessor struct{}

func (s BaseProcessor) Completion() {
	log.Info("base server's completion")
}
