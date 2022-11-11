package log

import (
	"errors"
	"testing"
)

func TestLogger(t *testing.T) {
	Info("Hello, World!")
	Warn("Hello, Warning!")
	Error("nil pointer dereference", ErrorField(errors.New("NPE")))
	Error("some error", ErrorStr(errors.New("what's wrong")))
	Infof("Today is %s, so happy!", "Friday")
	Debug("This log should not be displayed")
	SetDebug()
	Debug("This log should be displayed")
	SetLevel(ErrorLevel)
	Info("This log should not be displayed")
	Warn("This log should not be displayed")

	logger, _ := New(&Config{Encoding: "console"})
	ReplaceLogger(logger)
	SetLevel(InfoLevel)
	Infof("What's your name? My name is %s", "Chan")
	Error("illegal argument error", String("arg", "name"))
}
