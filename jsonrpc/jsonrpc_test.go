package jsonrpc

import (
	"testing"
	"time"

	"github.com/haobogu/lsframework/log"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"test1",
			"myresult",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewClient()
			if got := Call(); got != tt.want {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
			log.Info("sleeping")
			time.Sleep(1 * time.Second)
			if got := Call(); got != tt.want {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
			log.Info("end call")
		})
	}
}
