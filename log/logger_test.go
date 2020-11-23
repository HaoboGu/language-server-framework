package log

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Logger := NewLogger()
	Logger.Debug("This is a debug log")
	Logger.Infof("This is a %s log", "info")
}
