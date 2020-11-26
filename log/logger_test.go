package log

import (
	"testing"

	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	Debug("This is a debug ", "log")
	Infof("This is a %s log", "info")
	Info("This is a info level log")
	WithLogOption(zap.AddStacktrace(zap.DebugLevel)).Debug("This is a debug log")
	Debug("This is a debug log")
}
