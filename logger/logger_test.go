package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Info("Hello World!")
	Error("find a bug")
	Info("Hello World!")
}
